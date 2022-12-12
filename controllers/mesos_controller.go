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

package controllers

import (
	"context"
	"strings"

	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// MesosReconciler reconciles a Mesos object
type MesosReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=my.domain,resources=mesos,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=my.domain,resources=mesos/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=my.domain,resources=mesos/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Mesos object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile

func (r *MesosReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	reqLogger := log.Log
	reqLogger.Info("Reconciling Test")

	deployment := &v1.Deployment{}

	if err := r.Get(ctx, req.NamespacedName, deployment); err != nil {
		reqLogger.Info("Could not get deployment object: ")
		return ctrl.Result{}, nil
	}

	if !strings.Contains(deployment.Spec.Template.Spec.Containers[0].Image, "docker.io") {

		reqLogger.Info(deployment.Spec.Template.Spec.Containers[0].Image)
		deployment.Spec.Template.Spec.Containers[0].Image = "andererreposerver/" + deployment.Spec.Template.Spec.Containers[0].Image
		reqLogger.Info(deployment.Spec.Template.Spec.Containers[0].Image)

		if err := r.Update(ctx, deployment); err != nil {
			reqLogger.Info("Failed to remove finalizer for deployment")
			return ctrl.Result{}, err
		}
	}

	// your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MesosReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.Deployment{}).
		Complete(r)
}
