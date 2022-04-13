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

// BackupSpec defines the specification for a kahu backup.
type BackupSpec struct {

	ReclaimPolicy ReclaimPolicyType `json:"reclaimpolicy,omitempty"`

	MetadataLocation *metav1.BackupLocation `json:"metadataLocation,omitempty"`

	PreExecHook ResourceHookSpec `json:"preExecHook,omitempty"`

	PostExecHook ResourceHookSpec `json:"postExecHook,omitempty"`

	IncludedProviders []string `json:"includedProviders,omitempty"`

	ExcludeProviders []string `json:"excludedProviders,omitempty"`

	EnableMetadataBackup bool `json:"enableMetadataBackup,omitempty"`

	EnableVolumeBackup bool `json:"enableVolumeBackup,omitempty"`

	IncludedNamespaces []string `json:"includedNamespaces,omitempty"`

	ExcludedNamespaces []string `json:"excludedNamespaces,omitempty"`

	IncludedResources []string `json:"includedResources,omitempty"`

	ExcludedResources []string `json:"excludedResources,omitempty"`

    Label *metav1.LabelSelector `json:"label,omitempty"`

}

type ResourceHookSpec struct {

	Name string `json:"name"`

	IncludedNamespaces []string `json:"includedNamespaces,omitempty"`

	ExcludedNamespaces []string `json:"excludedNamespaces,omitempty"`

	IncludedResources []string `json:"includedResources,omitempty"`

	ExcludedResources []string `json:"excludedResources,omitempty"`

	LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty"`
}


type ReclaimPolicyType string

const (
	ReclaimPolicyDelete ReclaimPolicyType = "Delete"
	ReclaimPolicyRetain ReclaimPolicyType = "Retain"
)


type BackupPhase string

const (
	// BackupPhaseNew means the backup has been created but not
	// yet processed by the BackupController.
	BackupPhaseInit BackupPhase = "New"

	// BackupPhaseFailedValidation means the backup has failed
	// the controller's validations and therefore will not run.
	BackupPhaseFailedValidation BackupPhase = "FailedValidation"

	// BackupPhaseInProgress means the backup is currently executing.
	BackupPhaseInProgress BackupPhase = "InProgress"

	// BackupPhaseUploading means the backups of Kubernetes resources
	// and creation of snapshots was successful and snapshot data
	// is currently uploading.  The backup is not usable yet.
	BackupPhaseUploading BackupPhase = "Uploading"

	// BackupPhaseUploadingPartialFailure means the backup of Kubernetes
	// resources and creation of snapshots partially failed (final phase
	// will be PartiallyFailed) and snapshot data is currently uploading.
	// The backup is not usable yet.
	BackupPhaseUploadingPartialFailure BackupPhase = "UploadingPartialFailure"

	// BackupPhaseCompleted means the backup has run successfully without
	// errors.
	BackupPhaseCompleted BackupPhase = "Completed"

	// BackupPhasePartiallyFailed means the backup has run to completion
	// but encountered 1+ errors backing up individual items.
	BackupPhasePartiallyFailed BackupPhase = "PartiallyFailed"

	// BackupPhaseFailed means the backup ran but encountered an error that
	// prevented it from completing successfully.
	BackupPhaseFailed BackupPhase = "Failed"

	// BackupPhaseDeleting means the backup and all its associated data are being deleted.
	BackupPhaseDeleting BackupPhase = "Deleting"
)


type BackupCondition struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ResourceName string   `json:"resourceName,omitempty"`
	Status string `json:"status,omitempty"`
}

// BackupStatus defines the observed state of Backup
type BackupStatus struct {
	Phase BackupPhase `json:"phase,omitempty"`
    
	LastBackup *metav1.Time `json:"lastBackup,omitempty"`

	ValidationErrors []string `json:"validationErrors,omitempty"`

	Conditions []BackupCondition `json:"conditions,omitempty"`

}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Backup is the Schema for the backups API
type Backup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupSpec   `json:"spec,omitempty"`
	Status BackupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BackupList contains a list of Backup
type BackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Backup{}, &BackupList{})
}
