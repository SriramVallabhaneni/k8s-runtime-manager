/*
Copyright 2026.

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

package controller

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	runtimev1alpha1 "github.com/SriramVallabhaneni/ai-runtime-manager/operator/api/v1alpha1"
)

// AIModelDeploymentReconciler reconciles a AIModelDeployment object
type AIModelDeploymentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=runtime.airuntime.dev,resources=aimodeldeployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=runtime.airuntime.dev,resources=aimodeldeployments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=runtime.airuntime.dev,resources=aimodeldeployments/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the AIModelDeployment object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.24.1/pkg/reconcile
func (r *AIModelDeploymentReconciler) Reconcile(
	ctx context.Context,
	req ctrl.Request,
) (ctrl.Result, error) {
	logger := logf.FromContext(ctx)

	// Fetch the AIModelDeployment that triggered this reconciliation.
	aiModelDeployment := &runtimev1alpha1.AIModelDeployment{}

	err := r.Get(ctx, req.NamespacedName, aiModelDeployment)
	if err != nil {
		if apierrors.IsNotFound(err) {
			// The resource was deleted before this reconciliation ran.
			// There is nothing left for the controller to do.
			return ctrl.Result{}, nil
		}

		logger.Error(err, "unable to fetch AIModelDeployment")
		return ctrl.Result{}, err
	}

	// The child Deployment uses the same name and namespace as the
	// AIModelDeployment custom resource.
	deployment := &appsv1.Deployment{}

	err = r.Get(
		ctx,
		types.NamespacedName{
			Name:      aiModelDeployment.Name,
			Namespace: aiModelDeployment.Namespace,
		},
		deployment,
	)

	if err == nil {
		// The Deployment already exists.
		// For this first milestone, no update is necessary.
		logger.Info(
			"child Deployment already exists",
			"deployment",
			deployment.Name,
		)

		return ctrl.Result{}, nil
	}

	if !apierrors.IsNotFound(err) {
		logger.Error(err, "unable to fetch child Deployment")
		return ctrl.Result{}, err
	}

	// The Deployment does not exist, so build the desired child resource.
	deployment = r.deploymentForAIModel(aiModelDeployment)

	// Establish the ownership relationship:
	//
	// AIModelDeployment
	//        owns
	// Kubernetes Deployment
	//
	// This also allows Kubernetes garbage collection to delete the child
	// Deployment when the custom resource is deleted.
	if err := controllerutil.SetControllerReference(
		aiModelDeployment,
		deployment,
		r.Scheme,
	); err != nil {
		logger.Error(err, "unable to set owner reference on Deployment")
		return ctrl.Result{}, err
	}

	logger.Info(
		"creating child Deployment",
		"deployment",
		deployment.Name,
	)

	if err := r.Create(ctx, deployment); err != nil {
		logger.Error(err, "unable to create child Deployment")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *AIModelDeploymentReconciler) deploymentForAIModel(
	aiModelDeployment *runtimev1alpha1.AIModelDeployment,
) *appsv1.Deployment {
	labels := map[string]string{
		"app.kubernetes.io/name":       "ai-model-runtime",
		"app.kubernetes.io/instance":   aiModelDeployment.Name,
		"app.kubernetes.io/managed-by": "ai-runtime-operator",
	}

	replicas := aiModelDeployment.Spec.Replicas
	if replicas == 0 {
		replicas = 1
	}

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      aiModelDeployment.Name,
			Namespace: aiModelDeployment.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "runtime",
							Image: "nginx:1.27-alpine",
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *AIModelDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&runtimev1alpha1.AIModelDeployment{}).
		Owns(&appsv1.Deployment{}).
		Named("aimodeldeployment").
		Complete(r)
}
