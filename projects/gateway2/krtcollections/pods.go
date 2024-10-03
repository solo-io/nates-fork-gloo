package krtcollections

import (
	"maps"

	"istio.io/api/label"
	"istio.io/istio/pkg/kube"
	"istio.io/istio/pkg/kube/kclient"
	"istio.io/istio/pkg/kube/krt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

type nodeMetadata struct {
	name   string
	labels map[string]string
}

type PodLocality struct {
	Region  string
	Zone    string
	Subzone string
}

func (c nodeMetadata) ResourceName() string {
	return c.name
}
func (c nodeMetadata) Equals(in nodeMetadata) bool {
	return c.name == in.name && maps.Equal(c.labels, in.labels)
}

type LocalityPod struct {
	krt.Named
	Locality        PodLocality
	PodLabels       map[string]string
	AugmentedLabels map[string]string
}

func (c LocalityPod) Equals(in LocalityPod) bool {
	return c.Named == in.Named && c.Locality == in.Locality && maps.Equal(c.PodLabels, in.PodLabels) && maps.Equal(c.AugmentedLabels, in.AugmentedLabels)
}

func newNodeCollection(istioClient kube.Client) krt.Collection[nodeMetadata] {
	nodeClient := kclient.New[*corev1.Node](istioClient)
	nodes := krt.WrapClient(nodeClient, krt.WithName("Nodes"))
	return krt.NewCollection(nodes, func(kctx krt.HandlerContext, us *corev1.Node) *nodeMetadata {
		return &nodeMetadata{
			name:   us.Name,
			labels: us.Labels,
		}
	})
}

func NewPodsCollection(istioClient kube.Client) krt.Collection[LocalityPod] {
	podClient := kclient.New[*corev1.Pod](istioClient)
	pods := krt.WrapClient(podClient, krt.WithName("Pods"))
	nodes := newNodeCollection(istioClient)

	return krt.NewCollection(pods, augmentPodLabels(nodes))
}

func augmentPodLabels(nodes krt.Collection[nodeMetadata]) func(kctx krt.HandlerContext, pod *corev1.Pod) *LocalityPod {
	return func(kctx krt.HandlerContext, pod *corev1.Pod) *LocalityPod {
		labels := maps.Clone(pod.Labels)
		nodeName := pod.Spec.NodeName
		var l PodLocality
		if nodeName != "" {
			maybeNode := krt.FetchOne(kctx, nodes, krt.FilterObjectName(types.NamespacedName{
				Name: nodeName,
			}))
			if maybeNode != nil {
				node := *maybeNode
				nodeLabels := node.labels
				region := nodeLabels[corev1.LabelTopologyRegion]
				zone := nodeLabels[corev1.LabelTopologyZone]
				subzone := nodeLabels[label.TopologySubzone.Name]
				l = PodLocality{
					Region:  region,
					Zone:    zone,
					Subzone: subzone,
				}

				// augment labels
				labels[corev1.LabelTopologyRegion] = region
				labels[corev1.LabelTopologyZone] = zone
				labels[label.TopologySubzone.Name] = subzone
				//	labels[label.TopologyCluster.Name] = clusterID.String()
				//	labels[LabelHostname] = k8sNode
				//	labels[label.TopologyNetwork.Name] = networkID.String()
			}
		}

		return &LocalityPod{
			Named:           krt.NewNamed(pod),
			PodLabels:       pod.Labels,
			AugmentedLabels: labels,
			Locality:        l,
		}
	}

}
