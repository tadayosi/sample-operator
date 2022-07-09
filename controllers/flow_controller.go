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

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	samplev1alpha1 "github.com/tadayosi/sample-operator/api/v1alpha1"
)

// FlowReconciler reconciles a Flow object
type FlowReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=sample.tadayosi.github.io,resources=flows,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=sample.tadayosi.github.io,resources=flows/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=sample.tadayosi.github.io,resources=flows/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *FlowReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("Reconciling...")

	flow := &samplev1alpha1.Flow{}
	if err := r.Get(ctx, req.NamespacedName, flow); err != nil {
		return ctrl.Result{}, err
	}

	logger.Info("Flow:")
	j, err := json.Marshal(flow.Spec.Dsl)
	if err != nil {
		return ctrl.Result{}, err
	}
	m := make(map[string]interface{})
	if err = json.Unmarshal(j, &m); err != nil {
		return ctrl.Result{}, err
	}
	y, err := yaml.Marshal(m)
	if err != nil {
		return ctrl.Result{}, err
	}
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println(string(y))
	fmt.Println(strings.Repeat("=", 60))

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FlowReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&samplev1alpha1.Flow{}).
		Complete(r)
}
