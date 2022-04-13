/*
Copyright 2022.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RestoreSpec defines the desired state of Restore
type RestoreSpec struct {
	
	BackupName string `json:"backupName,omitempty"`

	IncludedNamespaces []string `json:"includedNamespaces,omitempty"`

	ExcludedNamespaces []string `json:"excludedNamespaces,omitempty"`

	IncludedResources []string `json:"includedResources,omitempty"`

	ExcludedResources []string `json:"excludedResources,omitempty"`

	NamespaceMapping map[string]string `json:"namespaceMapping,omitempty"`

    Label *metav1.LabelSelector `json:"label,omitempty"`

	RestorePVS bool `json:"restorepvs,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`
}



type RestorePhase string

const (
	// BackupPhaseNew means the backup has been created but not
	// yet processed by the BackupController.
	RestorePhaseInit RestorePhase = "New"

	// BackupPhaseFailedValidation means the backup has failed
	// the controller's validations and therefore will not run.
	RestorePhaseFailedValidation RestorePhase = "FailedValidation"

	// BackupPhaseInProgress means the backup is currently executing.
	RestorePhaseInProgress RestorePhase = "InProgress"

	// BackupPhaseUploading means the backups of Kubernetes resources
	// and creation of snapshots was successful and snapshot data
	// is currently uploading.  The backup is not usable yet.
	RestorePhaseUploading RestorePhase = "Uploading"

	// BackupPhaseUploadingPartialFailure means the backup of Kubernetes
	// resources and creation of snapshots partially failed (final phase
	// will be PartiallyFailed) and snapshot data is currently uploading.
	// The backup is not usable yet.
	RestorePhaseUploadingPartialFailure RestorePhase = "UploadingPartialFailure"

	// BackupPhaseCompleted means the backup has run successfully without
	// errors.
	RestorePhaseCompleted RestorePhase = "Completed"

	// BackupPhasePartiallyFailed means the backup has run to completion
	// but encountered 1+ errors backing up individual items.
	RestorePhasePartiallyFailed RestorePhase = "PartiallyFailed"

	// BackupPhaseFailed means the backup ran but encountered an error that
	// prevented it from completing successfully.
	RestorePhaseFailed RestorePhase = "Failed"

	// BackupPhaseDeleting means the backup and all its associated data are being deleted.
	RestorePhaseDeleting RestorePhase = "Deleting"
)

type RestoreProgress struct {
	TotalItems int `json:"totalItems,omitempty"`

	ItemsRestored int `json:"itemsRestored,omitempty"`
}

// RestoreStatus defines the observed state of Restore
type RestoreStatus struct {
	
	Phase RestorePhase `json:"phase,omitempty"`

	ValidationErrors []string `json:"validationErrors,omitempty"`

	Warnings int `json:"warnings,omitempty"`

	Errors int `json:"errors,omitempty"`

	FailureReason int `json:"failureReason,omitempty"`

	StartTimeStamp metav1.Time `json:"startTimeStamp,omitempty"`

	CompletionTimeStamp metav1.Time `json:"completionTimeStamp,omitempty"`

	Progress RestoreProgress `json:"progress,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Restore is the Schema for the restores API
type Restore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestoreSpec   `json:"spec,omitempty"`
	Status RestoreStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RestoreList contains a list of Restore
type RestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Restore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Restore{}, &RestoreList{})
}
