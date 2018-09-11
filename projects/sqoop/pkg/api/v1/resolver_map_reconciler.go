package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionResolverMapFunc func(original, desired *ResolverMap) (bool, error)

type ResolverMapReconciler interface {
	Reconcile(namespace string, desiredResources ResolverMapList, transition TransitionResolverMapFunc, opts clients.ListOpts) error
}

func resolverMapsToResources(list ResolverMapList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, resolverMap := range list {
		resourceList = append(resourceList, resolverMap)
	}
	return resourceList
}

func NewResolverMapReconciler(client ResolverMapClient) ResolverMapReconciler {
	return &resolverMapReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type resolverMapReconciler struct {
	base reconcile.Reconciler
}

func (r *resolverMapReconciler) Reconcile(namespace string, desiredResources ResolverMapList, transition TransitionResolverMapFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "resolverMap_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*ResolverMap), desired.(*ResolverMap))
		}
	}
	return r.base.Reconcile(namespace, resolverMapsToResources(desiredResources), transitionResources, opts)
}
