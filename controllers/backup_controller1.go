package controllers

import (
	kahuv1client "github.com/soda-cdm/kahu/pkg/generated/clientset/versioned/typed/kahu/v1"
	kahulister "github.com/soda-cdm/kahu/pkg/generated/listers/samplecontroller/v1alpha1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	k8sClient "sigs.k8s.io/controller-runtime/pkg/client"
)

type BackupController struct {
	k8sclientset      k8sClient.Interface
	kahuclientset     kahuv1client.BackupInterface
	kahuLister        kahulister.BackupLister
	deploymentsSynced cache.InformerSynced
	kahuSynced        cache.InformerSynced
	workqueue         workqueue.RateLimitingInterface
	recorder          record.EventRecorder
}


type NewBackupController struct{
	k8sclientset kubernetes.Interface,
	kahuclientset clientset.Interface,
	deploymentInformer appsinformers.DeploymentInformer,
	kahuInformer informers.FooInformer) *Controller {

	klog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset:     kubeclientset,
		kahuclientset:   kahuclientset,
		deploymentsLister: deploymentInformer.Lister(),
		deploymentsSynced: deploymentInformer.Informer().HasSynced,
		kahuLister:        kahuInformer.Lister(),
		kahuSynced:        kahuInformer.Informer().HasSynced,
		workqueue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Foos"),
		recorder:          recorder,
	}

	klog.Info("Setting up event handlers")
	// Set up an event handler for when Foo resources change
	fooInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueKahu,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueKahu(new)
		},
		DeleteFunc: controller.enqueueKahu,

	})
	deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.handleObject,
		UpdateFunc: func(old, new interface{}) {
			newDepl := new.(*appsv1.Deployment)
			oldDepl := old.(*appsv1.Deployment)
			if newDepl.ResourceVersion == oldDepl.ResourceVersion {
				return
			}
			controller.handleObject(new)
		},
		DeleteFunc: controller.handleObject,
	})

	return controller
}

func (c *BackupController) Run(workers int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()

	klog.Info("Starting Backup controller")

	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.deploymentsSynced, c.kahuSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	klog.Info("Starting workers")
	for i := 0; i < workers; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	klog.Info("Started workers")
	<-stopCh
	klog.Info("Shutting down workers")

	return nil
}

func (c *BackupController) runWorker() {
	for c.processNextWorkItem() {
	}
}

func (c *genericController) processNextWorkItem() bool {
}