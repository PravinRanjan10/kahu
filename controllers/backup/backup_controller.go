/*
Copyright 2022 The SODA Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package backup

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	kahuv1beta1 "github.com/soda-cdm/kahu/apis/kahu/v1beta1"
	pkgbackup "github.com/soda-cdm/kahu/controllers"
	"google.golang.org/grpc"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"

	"github.com/soda-cdm/kahu/apis/kahu/v1beta1"
	kahuclientset "github.com/soda-cdm/kahu/client/clientset/versioned"
	kahuv1client "github.com/soda-cdm/kahu/client/clientset/versioned/typed/kahu/v1beta1"
	kinf "github.com/soda-cdm/kahu/client/informers/externalversions/kahu/v1beta1"
	kahulister "github.com/soda-cdm/kahu/client/listers/kahu/v1beta1"
	"github.com/soda-cdm/kahu/controllers"
	"github.com/soda-cdm/kahu/controllers/backup/cmd/options"
	metaservice "github.com/soda-cdm/kahu/providerframework/meta_service/lib/go"
	utils "github.com/soda-cdm/kahu/utils"
	"k8s.io/client-go/dynamic"
)

var (
	grpcServer     *grpc.Server
	grpcConnection *grpc.ClientConn
	metaClient     metaservice.MetaServiceClient
)

type Controller struct {
	*controllers.BaseController
	client       kubernetes.Interface
	klient       kahuclientset.Interface
	backupSynced []cache.InformerSynced
	kLister      kahulister.BackupLister
	config       *restclient.Config
	bkpClient    kahuv1client.BackupsGetter
	bkpLocClient kahuv1client.BackupLocationsGetter
	flags        *options.BackupControllerFlags
}

func NewController(klient kahuclientset.Interface,
	backupInformer kinf.BackupInformer,
	config *restclient.Config,
	bkpClient kahuv1client.BackupsGetter,
	bkpLocClient kahuv1client.BackupLocationsGetter,
	flags *options.BackupControllerFlags) *Controller {

	c := &Controller{
		BaseController: controllers.NewBaseController(controllers.Backup),
		klient:         klient,
		kLister:        backupInformer.Lister(),
		config:         config,
		bkpClient:      bkpClient,
		bkpLocClient:   bkpLocClient,
		flags:          flags,
	}

	backupInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    c.handleAdd,
			DeleteFunc: c.handleDel,
		},
	)

	return c
}

func (c *Controller) Run(ch chan struct{}) error {
	// only want to log about cache sync waiters if there are any
	if len(c.backupSynced) > 0 {
		c.Logger.Infoln("Waiting for caches to sync")
		if !cache.WaitForCacheSync(context.Background().Done(), c.backupSynced...) {
			return errors.New("timed out waiting for caches to sync")
		}
		c.Logger.Infoln("Caches are synced")
	}

	if ok := cache.WaitForCacheSync(ch, c.backupSynced...); !ok {
		c.Logger.Infoln("cache was not sycned")
	}

	go wait.Until(c.worker, time.Second, ch)

	<-ch
	return nil
}

func (c *Controller) worker() {
	for c.processNextItem() {
	}
}

type restoreContext struct {
	namespaceClient corev1.NamespaceInterface
}

func (c *Controller) processNextItem() bool {

	item, shutDown := c.Workqueue.Get()
	if shutDown {
		return false
	}

	defer c.Workqueue.Forget(item)

	key, err := cache.MetaNamespaceKeyFunc(item)
	if err != nil {
		c.Logger.Errorf("error %s calling Namespace key func on cache for item", err.Error())
		return false
	}

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		c.Logger.Errorf("splitting key into namespace and name, error %s\n", err.Error())
		return false
	}

	backupObj, err := c.kLister.Backups(namespace).Get(name)
	if err != nil {
		c.Logger.Errorf("error %s, Getting the backup resource from lister", err.Error())

		if apierrors.IsNotFound(err) {
			c.Logger.Debugf("backup %s not found", name)
		}
		return false
	}

	// check if Metadatalocation is empty
	if backupObj.Spec.MetadataLocation == nil {
		c.Logger.Errorf("backup location can not be empty")
		backupObj.Status.Phase = kahuv1beta1.BackupPhaseFailed
		c.updateStatus(backupObj, backupObj, c.bkpClient, backupObj.Status.Phase)
		return false
	}

	// Validate the Metadatalocation
	backupProvider := backupObj.Spec.MetadataLocation.Spec.ProviderName
	c.Logger.Infof("preparing backup for provider %s: ", backupProvider)
	_, err = c.bkpLocClient.BackupLocations("default").Get(context.Background(), backupProvider, metav1.GetOptions{})
	if err != nil {
		c.Logger.Errorf("failed to validate backup location because %s", err)
		backupObj.Status.Phase = kahuv1beta1.BackupPhaseFailedValidation
		c.updateStatus(backupObj, backupObj, c.bkpClient, backupObj.Status.Phase)
		return false
	}

	c.Logger.Infof("Preparing backup request for Provider:%s", backupObj.Spec.MetadataLocation.Spec.ProviderName)
	request := c.prepareBackupRequest(backupObj)

	if len(request.Status.ValidationErrors) > 0 {
		request.Status.Phase = kahuv1beta1.BackupPhaseFailedValidation
	} else {
		request.Status.StartTimestamp = &metav1.Time{Time: time.Now()}
		request.Status.Phase = kahuv1beta1.BackupPhaseInProgress
	}
	request.Status.StartTimestamp = &metav1.Time{Time: time.Now()}
	c.updateStatus(backupObj, request.Backup, c.bkpClient, request.Status.Phase)

	// prepare and run backup
	c.runBackup(backupObj)

	return true
}

type kubernetesResource struct {
	groupResource         schema.GroupResource
	preferredGVR          schema.GroupVersionResource
	namespace, name, path string
}

// var KindList map[string]int
var kindList = make(map[string]groupResouceVersion)

func (c *Controller) runBackup(backup *v1beta1.Backup) error {
	c.Logger.WithField(controllers.Backup, utils.NamespaceAndName(backup)).Info("Setting up backup log")

	c.Logger.Infof("runBackup: starting backup for obj:%+v", backup)

	var resourcesList []*kubernetesResource

	// for _, group := range c.discoveryHelper.Resources() {
	k8sClinet, err := kubernetes.NewForConfig(c.config)
	if err != nil {
		c.Logger.Errorf("unable to get k8s client:%s", err)
		return err
	}

	_, resource, _ := k8sClinet.ServerGroupsAndResources()
	for _, group := range resource {
		c.Logger.Infof("group *****************:%+v", group)

		groupItems, err := c.getGroupItems(group)
		c.Logger.Infof("groupItems *****************:%+v", groupItems)

		if err != nil {
			c.Logger.WithError(err).WithField("apiGroup", group.String()).Error("Error collecting resources from API group")
			continue
		}

		c.Logger.Infof("everything to be backuped")
		resourcesList = append(resourcesList, groupItems...)

	}
	c.Logger.Infof("resources $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$:%+v", resourcesList)
	c.Logger.Infof("kindList ===================================>\n:%+v", kindList)
	c.Logger.Infof("kindList ===================================>\n:%+v", len(kindList))

	for _, kind := range kindList {
		c.fun(kind.group, kind.version, kind.resourceName)
	}

	return nil
}

// getGroupItems collects all relevant items from a single API group.
func (c *Controller) getGroupItems(group *metav1.APIResourceList) ([]*kubernetesResource, error) {
	c.Logger.WithField("group", group.GroupVersion)

	c.Logger.Infof("Getting items for group")

	// Parse so we can check if this is the core group
	gv, err := schema.ParseGroupVersion(group.GroupVersion)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing GroupVersion %q", group.GroupVersion)
	}
	// c.log.Infoln("the gv------------------------------:", gv)
	if gv.Group == "" {
		// This is the core group, so make sure we process in the following order: pods, pvcs, pvs,
		// everything else.
		sortCoreGroup(group)
	}

	var items []*kubernetesResource
	for _, resource := range group.APIResources {
		c.Logger.Infof("getGroupItems->resource:%+v", resource)
		// kindList = append(kindList, resource.Kind)

		// val, ok := kindList[resource.Kind]
		// c.Logger.Infoln("resource.Kind, val and ok", resource.Kind, val, ok)

		// if !ok {
		// 	c.Logger.Infoln("resource.Kind, val and ok", resource.Kind, val, ok)
		// 	kindList[resource.Kind] =
		// 	i = i + 1
		// }
		resourceItems, err := c.getResourceItems(gv, resource)
		if err != nil {
			log.WithError(err).WithField("resource", resource.String()).Error("Error getting items for resource")
			continue
		}

		items = append(items, resourceItems...)
	}

	return items, nil
}

func (c *Controller) fun(group, version, resource string) bool {

	dynamicClient, err := dynamic.NewForConfig(c.config)
	if err != nil {
		fmt.Printf("error creating dynamic client: %v\n", err)
		os.Exit(1)
	}

	gvr := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}

	c.Logger.Infof("group:%s, version:%s, resource:%s", gvr.Group, gvr.Version, gvr.Resource)
	objectsList, err := dynamicClient.Resource(gvr).Namespace("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error getting %s, %v\n", resource, err)
		// os.Exit(1)
	}

	// for _, item := range objectsList.Items {
	// 	resource_data, err := k8sClinet.CoreV1().Pods(ns).Get(context.TODO(), item.Name, metav1.GetOptions{})
	// 	if err != nil {
	// 		c.Logger.Errorf("Unable to get resource content: %s", err)
	// 	}
	// }
	if objectsList != nil {
		for _, d := range objectsList.Items {
			replicas, found, err := unstructured.NestedInt64(d.Object, "spec", "replicas")
			if err != nil || !found {
				fmt.Printf("Replicas not found for deployment %s: error=%s", d.GetName(), err)
				continue
			}
			fmt.Printf(" * %s (%d replicas)\n", d.GetName(), replicas)
		}
	}

	resourceObjects, err := meta.ExtractList(objectsList)
	if err != nil {
		return false
	}

	for _, o := range resourceObjects {
		c.Logger.Infoln("************************************************")
		runtimeObject, ok := o.(runtime.Unstructured)
		if !ok {
			c.Logger.Errorf("error casting object: %v", o)
			return false
		}
		// c.Logger.Infoln("THE runtimeObject*************\n:", runtimeObject)

		metadata, err := meta.Accessor(runtimeObject)
		if err != nil {
			return false
		}
		file, _ := json.MarshalIndent(metadata, "", " ")
		// c.Logger.Infoln("THE metadata********************\n:", metadata)
		filename := resource + "_" + metadata.GetName() + ".json"
		c.Logger.Infoln("filename:", filename)
		_ = ioutil.WriteFile(filename, file, 0644)
		// WriteToFile(metadata)

	}
	return true
}

func WriteToFile(metadata metav1.Object) (string, error) {
	f, err := ioutil.TempFile("/root/dumpdata", "")
	if err != nil {
		return "", errors.Wrap(err, "error creating temp file")
	}
	defer f.Close()

	jsonBytes, err := json.Marshal(metadata)
	if err != nil {
		return "", errors.Wrap(err, "error converting item to JSON")
	}

	if _, err := f.Write(jsonBytes); err != nil {
		return "", errors.Wrap(err, "error writing JSON to file")
	}

	if err := f.Close(); err != nil {
		return "", errors.Wrap(err, "error closing file")
	}

	return f.Name(), nil
}

type groupResouceVersion struct {
	resourceName string
	version      string
	group        string
}

// getResourceItems collects all relevant items for a given group-version-resource.
func (r *Controller) getResourceItems(gv schema.GroupVersion, resource metav1.APIResource) ([]*kubernetesResource, error) {
	log.WithField("resource", resource.Name)

	log.Info("Getting items for resource")

	var (
		gvr           = gv.WithResource(resource.Name)
		gr            = gvr.GroupResource()
		clusterScoped = !resource.Namespaced
	)

	// gvr := schema.GroupVersionResource{
	// 	Group:    "apps",
	// 	Version:  "v1",
	// 	Resource: "Deployments",
	// }
	log.Infoln("gvr:", gvr)
	log.Infoln("gr:", gr)
	log.Infoln("gv:", gv)
	log.Infoln("clusterScoped:", clusterScoped)

	_, ok := kindList[resource.Kind]
	if !ok {
		gvr1 := groupResouceVersion{
			resourceName: gvr.Resource,
			version:      gvr.Version,
			group:        gvr.Group,
		}
		kindList[resource.Kind] = gvr1
	}

	// orders := getOrderedResourcesForType(log, nil, resource.Name)
	// // Getting the preferred group version of this resource
	// preferredGVR, _, err := r.discoveryHelper.ResourceFor(gr.WithVersion(""))
	// if err != nil {
	// 	return nil, errors.WithStack(err)
	// }
	// r.log.Infoln("the preferredGVR============>:", preferredGVR)

	// // If the resource we are backing up is NOT namespaces, and it is cluster-scoped, check to see if
	// // we should include it based on the IncludeClusterResources setting.
	// // if gr != Namespaces && clusterScoped {
	// // 	if !r.backupRequest.NamespaceIncludesExcludes.IncludeEverything() {
	// // 		// when IncludeClusterResources == nil (auto), only directly
	// // 		// back up cluster-scoped resources if we're doing a full-cluster
	// // 		// (all namespaces) backup. Note that in the case of a subset of
	// // 		// namespaces being backed up, some related cluster-scoped resources
	// // 		// may still be backed up if triggered by a custom action (e.g. PVC->PV).
	// // 		// If we're processing namespaces themselves, we will not skip here, they may be
	// // 		// filtered out later.
	// // 		log.Info("Skipping resource because it's cluster-scoped and only specific namespaces are included in the backup")
	// // 		return nil, nil
	// // 	}
	// // }

	// // if !r.backupRequest.ResourceIncludesExcludes.ShouldInclude(gr.String()) {
	// // 	log.Infof("Skipping resource because it's excluded")
	// // 	return nil, nil
	// // }

	// if cohabitator, found := r.cohabitatingResources[resource.Name]; found {
	// 	if gv.Group == cohabitator.groupResource1.Group || gv.Group == cohabitator.groupResource2.Group {
	// 		if cohabitator.seen {
	// 			log.WithFields(
	// 				logrus.Fields{
	// 					"cohabitatingResource1": cohabitator.groupResource1.String(),
	// 					"cohabitatingResource2": cohabitator.groupResource2.String(),
	// 				},
	// 			).Infof("Skipping resource because it cohabitates and we've already processed it")
	// 			return nil, nil
	// 		}
	// 		cohabitator.seen = true
	// 	}
	// }

	// namespacesToList := getNamespacesToList(r.backupRequest.NamespaceIncludesExcludes)

	// namespacesToList = append(namespacesToList, "kube-system")
	// // r.log.Infoln("the namespaceTo List:================>:", namespacesToList)

	// // Check if we're backing up namespaces, and only certain ones
	// if gr == Namespaces && namespacesToList[0] != "" {
	// 	resourceClient, err := r.dynamicFactory.ClientForGroupVersionResource(gv, resource, "")
	// 	// r.log.Infoln("the resourceClient:================>:", resourceClient)
	// 	if err != nil {
	// 		log.WithError(err).Error("Error getting dynamic client")
	// 	} else {
	// 		var labelSelector labels.Selector
	// 		if r.backupRequest.Spec.Label != nil {
	// 			labelSelector, err = metav1.LabelSelectorAsSelector(r.backupRequest.Spec.Label)
	// 			if err != nil {
	// 				// This should never happen...
	// 				return nil, errors.Wrap(err, "invalid label selector")
	// 			}
	// 		}

	// 		var items []*kubernetesResource
	// 		namespacesToList = append(namespacesToList, "kube-system")
	// 		for _, ns := range namespacesToList {
	// 			log = log.WithField("namespace", ns)
	// 			// log.Info("Getting namespace:", ns)
	// 			unstructured, err := resourceClient.Get(ns, metav1.GetOptions{})
	// 			if err != nil {
	// 				log.WithError(errors.WithStack(err)).Error("Error getting namespace")
	// 				continue
	// 			}

	// 			labels := labels.Set(unstructured.GetLabels())
	// 			if labelSelector != nil && !labelSelector.Matches(labels) {
	// 				log.Info("Skipping namespace because it does not match the backup's label selector")
	// 				continue
	// 			}

	// 			path, err := r.writeToFile(unstructured)
	// 			if err != nil {
	// 				log.WithError(err).Error("Error writing item to file")
	// 				continue
	// 			}

	// 			items = append(items, &kubernetesResource{
	// 				groupResource: gr,
	// 				// preferredGVR:  preferredGVR,
	// 				name: ns,
	// 				path: path,
	// 			})
	// 		}
	// 		// r.log.Infoln("the items:######################################>:", items)
	// 		return items, nil
	// 	}
	// }

	// // If we get here, we're backing up something other than namespaces
	// if clusterScoped {
	// 	namespacesToList = []string{""}
	// }

	var items []*kubernetesResource

	// for _, namespace := range namespacesToList {
	// 	// List items from Kubernetes API
	// 	// log = log.WithField("namespace", namespace)

	// resourceClient, err := r.dynamicFactory.ClientForGroupVersionResource(gv, resource, namespace)
	// if err != nil {
	// 	log.WithError(err).Error("Error getting dynamic client")
	// 	continue
	// }

	// 	var labelSelector string
	// 	if selector := r.backupRequest.Spec.Label; selector != nil {
	// 		labelSelector = metav1.FormatLabelSelector(selector)
	// 	}

	// 	// log.Info("Listing items")
	// 	unstructuredItems := make([]unstructured.Unstructured, 0)

	// 	// If limit is not positive, do not use paging. Instead, request all items at once
	// 	unstructuredList, err := resourceClient.List(metav1.ListOptions{LabelSelector: labelSelector})
	// 	if err != nil {
	// 		log.WithError(errors.WithStack(err)).Error("Error listing items")
	// 		continue
	// 	}
	// 	unstructuredItems = append(unstructuredItems, unstructuredList.Items...)

	// 	log.Infof("Retrieved %d items", len(unstructuredItems))

	// 	for i, uunstructuredItem := range unstructuredItems {
	// 		log.Infof("Retrieved uunstructuredItem no: %d and %d item", i, uunstructuredItem.GetName())
	// 	}

	// 	//Collect items in included Namespaces
	// 	for i := range unstructuredItems {
	// 		item := &unstructuredItems[i]

	// 		// if gr == Namespaces && !r.backupRequest.NamespaceIncludesExcludes.ShouldInclude(item.GetName()) {
	// 		// 	log.WithField("name", item.GetName()).Info("Skipping namespace because it's excluded")
	// 		// 	continue
	// 		// }

	// 		path, err := r.writeToFile(item)
	// 		if err != nil {
	// 			log.WithError(err).Error("Error writing item to file")
	// 			continue
	// 		}

	// 		items = append(items, &kubernetesResource{
	// 			groupResource: gr,
	// 			preferredGVR:  preferredGVR,
	// 			namespace:     item.GetNamespace(),
	// 			name:          item.GetName(),
	// 			path:          path,
	// 		})
	// 	}
	// }

	// items = sortResourcesByOrder(r.log, items, orders)
	// r.log.Infoln("the items:********************************************>:", items)

	return items, nil
}

const (
	pod = iota
	pvc
	pv
	other
)

// coreGroupResourcePriority returns the relative priority of the resource, in the following order:
// pods, pvcs, pvs, everything else.
func coreGroupResourcePriority(resource string) int {
	switch strings.ToLower(resource) {
	case "pods":
		return pod
	case "persistentvolumeclaims":
		return pvc
	case "persistentvolumes":
		return pv
	}

	return other
}

// sortCoreGroup sorts the core API group.
func sortCoreGroup(group *metav1.APIResourceList) {
	sort.SliceStable(group.APIResources, func(i, j int) bool {
		return coreGroupResourcePriority(group.APIResources[i].Name) < coreGroupResourcePriority(group.APIResources[j].Name)
	})
}

func (c *Controller) updateStatus(bkp, updated *v1beta1.Backup, client kahuv1client.BackupsGetter, phase kahuv1beta1.BackupPhase) {
	backup, err := client.Backups(updated.Namespace).Get(context.Background(), bkp.Name, metav1.GetOptions{})
	if err != nil {
		c.Logger.Errorf("failed to get backup for updating status :%+s", err)
		return
	}

	backup.Status.Phase = phase
	backup.Status.ValidationErrors = updated.Status.ValidationErrors
	_, err = client.Backups(updated.Namespace).UpdateStatus(context.Background(), backup, metav1.UpdateOptions{})
	if err != nil {
		c.Logger.Errorf("failed to update backup status :%+s", err)
	}

	return
}
func (c *Controller) prepareBackupRequest(backup *kahuv1beta1.Backup) *pkgbackup.Request {
	request := &pkgbackup.Request{
		Backup: backup.DeepCopy(),
	}

	if request.Annotations == nil {
		request.Annotations = make(map[string]string)
	}

	// add the storage location as a label for easy filtering later.
	if request.Labels == nil {
		request.Labels = make(map[string]string)
	}

	// validate the included/excluded resources
	for _, err := range utils.ValidateIncludesExcludes(request.Spec.IncludedResources, request.Spec.ExcludedResources) {
		request.Status.ValidationErrors = append(request.Status.ValidationErrors, fmt.Sprintf("Invalid included/excluded resource lists: %v", err))
	}

	// validate the included/excluded namespaces
	for _, err := range utils.ValidateNamespaceIncludesExcludes(request.Spec.IncludedNamespaces, request.Spec.ExcludedNamespaces) {
		request.Status.ValidationErrors = append(request.Status.ValidationErrors, fmt.Sprintf("Invalid included/excluded namespace lists: %v", err))
	}

	// c.Logger.Infof("THe request:%+v", request.Backup)
	request.Status.Phase = v1beta1.BackupPhaseInit

	c.Logger.Infoln("validation done :", request.Status.ValidationErrors)
	// c.Logger.Infof("THe request after:%+v", request.Backup)
	return request
}

// NamespaceAndName returns a string in the format <namespace>/<name>
func NamespaceAndName(objMeta metav1.Object) string {
	if objMeta.GetNamespace() == "" {
		return objMeta.GetName()
	}
	return fmt.Sprintf("%s/%s", objMeta.GetNamespace(), objMeta.GetName())
}

func (c *Controller) handleAdd(obj interface{}) {
	backup := obj.(*v1beta1.Backup)

	switch backup.Status.Phase {
	case "", v1beta1.BackupPhaseInit:
	default:
		c.Logger.WithFields(logrus.Fields{
			"backup": utils.NamespaceAndName(backup),
			"phase":  backup.Status.Phase,
		}).Infof("Backup: %s is not New, so will not be processed", backup.Name)
		return
	}
	c.Workqueue.Add(obj)
}

func (c *Controller) handleDel(obj interface{}) {
	c.Workqueue.Add(obj)
}
