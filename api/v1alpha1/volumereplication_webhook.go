/*
Copyright 2021.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var volumereplicationlog = logf.Log.WithName("volumereplication-resource")

func (r *VolumeReplication) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-replication-storage-openshift-io-v1alpha1-volumereplication,mutating=true,failurePolicy=fail,sideEffects=None,groups=replication.storage.openshift.io,resources=volumereplications,verbs=create;update,versions=v1alpha1,name=mvolumereplication.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Defaulter = &VolumeReplication{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *VolumeReplication) Default() {
	volumereplicationlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-replication-storage-openshift-io-v1alpha1-volumereplication,mutating=false,failurePolicy=fail,sideEffects=None,groups=replication.storage.openshift.io,resources=volumereplications,verbs=create;update,versions=v1alpha1,name=vvolumereplication.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &VolumeReplication{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *VolumeReplication) ValidateCreate() error {
	volumereplicationlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *VolumeReplication) ValidateUpdate(old runtime.Object) error {
	volumereplicationlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *VolumeReplication) ValidateDelete() error {
	volumereplicationlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}