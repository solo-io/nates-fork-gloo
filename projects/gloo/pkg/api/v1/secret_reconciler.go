package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionSecretFunc func(original, desired *Secret) (bool, error)

type SecretReconciler interface {
	Reconcile(namespace string, desiredResources SecretList, transition TransitionSecretFunc, opts clients.ListOpts) error
}

func secretsToResources(list SecretList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, secret := range list {
		resourceList = append(resourceList, secret)
	}
	return resourceList
}

func NewSecretReconciler(client SecretClient) SecretReconciler {
	return &secretReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type secretReconciler struct {
	base reconcile.Reconciler
}

func (r *secretReconciler) Reconcile(namespace string, desiredResources SecretList, transition TransitionSecretFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "secret_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*Secret), desired.(*Secret))
		}
	}
	return r.base.Reconcile(namespace, secretsToResources(desiredResources), transitionResources, opts)
}
