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
func NewArtifact(namespace, name string) *Artifact {
	return &Artifact{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (r *Artifact) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *Artifact) SetData(data map[string]string) {
	r.Data = data
}

type ArtifactList []*Artifact
type ArtifactsByNamespace map[string]ArtifactList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list ArtifactList) Find(namespace, name string) (*Artifact, error) {
	for _, artifact := range list {
		if artifact.Metadata.Name == name {
			if namespace == "" || artifact.Metadata.Namespace == namespace {
				return artifact, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find artifact %v.%v", namespace, name)
}

func (list ArtifactList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, artifact := range list {
		ress = append(ress, artifact)
	}
	return ress
}

func (list ArtifactList) Names() []string {
	var names []string
	for _, artifact := range list {
		names = append(names, artifact.Metadata.Name)
	}
	return names
}

func (list ArtifactList) NamespacesDotNames() []string {
	var names []string
	for _, artifact := range list {
		names = append(names, artifact.Metadata.Namespace+"."+artifact.Metadata.Name)
	}
	return names
}

func (list ArtifactList) Sort() {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
}

func (list ArtifactList) Clone() ArtifactList {
	var artifactList ArtifactList
	for _, artifact := range list {
		artifactList = append(artifactList, proto.Clone(artifact).(*Artifact))
	}
	return artifactList
}

func (list ArtifactList) ByNamespace() ArtifactsByNamespace {
	byNamespace := make(ArtifactsByNamespace)
	for _, artifact := range list {
		byNamespace.Add(artifact)
	}
	return byNamespace
}

func (byNamespace ArtifactsByNamespace) Add(artifact ...*Artifact) {
	for _, item := range artifact {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace ArtifactsByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace ArtifactsByNamespace) List() ArtifactList {
	var list ArtifactList
	for _, artifactList := range byNamespace {
		list = append(list, artifactList...)
	}
	list.Sort()
	return list
}

func (byNamespace ArtifactsByNamespace) Clone() ArtifactsByNamespace {
	return byNamespace.List().Clone().ByNamespace()
}

var _ resources.Resource = &Artifact{}

// Kubernetes Adapter for Artifact

func (o *Artifact) GetObjectKind() schema.ObjectKind {
	t := ArtifactCrd.TypeMeta()
	return &t
}

func (o *Artifact) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Artifact)
}

var ArtifactCrd = crd.NewCrd("gloo.solo.io",
	"artifacts",
	"gloo.solo.io",
	"v1",
	"Artifact",
	"art",
	&Artifact{})
