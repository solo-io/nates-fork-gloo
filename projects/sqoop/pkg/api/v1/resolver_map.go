package v1

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: modify as needed to populate additional fields
func NewResolverMap(namespace, name string) *ResolverMap {
	return &ResolverMap{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (r *ResolverMap) SetStatus(status core.Status) {
	r.Status = status
}

func (r *ResolverMap) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

type ResolverMapList []*ResolverMap
type ResolverMapsByNamespace map[string]ResolverMapList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list ResolverMapList) Find(namespace, name string) (*ResolverMap, error) {
	for _, resolverMap := range list {
		if resolverMap.Metadata.Name == name {
			if namespace == "" || resolverMap.Metadata.Namespace == namespace {
				return resolverMap, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find resolverMap %v.%v", namespace, name)
}

func (list ResolverMapList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, resolverMap := range list {
		ress = append(ress, resolverMap)
	}
	return ress
}

func (list ResolverMapList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, resolverMap := range list {
		ress = append(ress, resolverMap)
	}
	return ress
}

func (list ResolverMapList) Names() []string {
	var names []string
	for _, resolverMap := range list {
		names = append(names, resolverMap.Metadata.Name)
	}
	return names
}

func (list ResolverMapList) NamespacesDotNames() []string {
	var names []string
	for _, resolverMap := range list {
		names = append(names, resolverMap.Metadata.Namespace+"."+resolverMap.Metadata.Name)
	}
	return names
}

func (list ResolverMapList) Sort() {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
}

func (list ResolverMapList) Clone() ResolverMapList {
	var resolverMapList ResolverMapList
	for _, resolverMap := range list {
		resolverMapList = append(resolverMapList, proto.Clone(resolverMap).(*ResolverMap))
	}
	return resolverMapList
}

func (list ResolverMapList) ByNamespace() ResolverMapsByNamespace {
	byNamespace := make(ResolverMapsByNamespace)
	for _, resolverMap := range list {
		byNamespace.Add(resolverMap)
	}
	return byNamespace
}

func (byNamespace ResolverMapsByNamespace) Add(resolverMap ...*ResolverMap) {
	for _, item := range resolverMap {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace ResolverMapsByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace ResolverMapsByNamespace) List() ResolverMapList {
	var list ResolverMapList
	for _, resolverMapList := range byNamespace {
		list = append(list, resolverMapList...)
	}
	list.Sort()
	return list
}

func (byNamespace ResolverMapsByNamespace) Clone() ResolverMapsByNamespace {
	return byNamespace.List().Clone().ByNamespace()
}

var _ resources.Resource = &ResolverMap{}

// Kubernetes Adapter for ResolverMap

func (o *ResolverMap) GetObjectKind() schema.ObjectKind {
	t := ResolverMapCrd.TypeMeta()
	return &t
}

func (o *ResolverMap) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*ResolverMap)
}

var ResolverMapCrd = crd.NewCrd("sqoop.solo.io",
	"resolvermaps",
	"sqoop.solo.io",
	"v1",
	"ResolverMap",
	"rm",
	&ResolverMap{})
