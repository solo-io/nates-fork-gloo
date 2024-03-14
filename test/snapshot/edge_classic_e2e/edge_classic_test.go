package edge_classic_e2e

import (
	"fmt"
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	gloodefaults "github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/gloo/test/kube2e"
	"github.com/solo-io/gloo/test/kube2e/helper"
	"github.com/solo-io/gloo/test/snapshot"
	"github.com/solo-io/gloo/test/snapshot/testcases"
	"github.com/solo-io/gloo/test/snapshot/utils"
	"github.com/solo-io/go-utils/testutils"
	v1 "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"
	gloov1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	"github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/core/matchers"
	"github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/kubernetes"
	gloocore "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	skhelpers "github.com/solo-io/solo-kit/test/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Gloo Edge Classic", func() {

	var (
		inputs  []client.Object
		kubeCtx string

		testHelper        *helper.SoloTestHelper
		resourceClientset *kube2e.KubeResourceClientSet
		kubeClient        client.Client
		runner            snapshot.TestRunner
	)

	BeforeEach(func() {
		// Clear inputs before each run.
		inputs = []client.Object{}
	})

	JustAfterEach(func() {
		defer ctxCancel()

		// Note to devs:  set NO_CLEANUP to 'all' or 'failed' to skip cleanup, for the sake of
		// debugging or otherwise examining state after a test.
		if utils.ShouldSkipCleanup() {
			fmt.Printf("Not cleaning up")
			return // Exit without cleaning up
		}
		Expect(runner.Cleanup(ctx)).To(Succeed())
	})

	When("Happy Path", func() {

		var err error
		testHelper, err = kube2e.GetTestHelper(ctx, gloodefaults.GlooSystem)
		Expect(err).NotTo(HaveOccurred())
		skhelpers.RegisterPreFailHandler(helpers.StandardGlooDumpOnFail(GinkgoWriter, testHelper.InstallNamespace))

		clusterName := os.Getenv("CLUSTER_NAME")
		kubeCtx = fmt.Sprintf("kind-%s", clusterName)

		resourceClientset, err = kube2e.NewDefaultKubeResourceClientSet(ctx)
		Expect(err).NotTo(HaveOccurred(), "can create kube resource client set")

		kubeClient, err = utils.GetClient(kubeCtx)
		Expect(err).NotTo(HaveOccurred(), "can create client")

		runner = snapshot.TestRunner{
			Name:             "classic-apis",
			ResultsByGateway: map[types.NamespacedName]snapshot.ExpectedTestResult{},
			ClientSet:        resourceClientset,
			Client:           kubeClient,
		}

		customSetupAssertions := func() {
			err := testutils.WaitPodsRunning(ctx, time.Second, testHelper.InstallNamespace, "gloo=gateway-proxy")
			Expect(err).NotTo(HaveOccurred())
		}

		// TODO: create builder for solo-apis VirtualService and Upstream
		inputs = []client.Object{
			// NOTE: this is the solo-apis VirtualService
			&gloov1.Upstream{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "httpbin-htppbin-8000",
					Namespace: gloodefaults.GlooSystem,
				},
				Spec: gloov1.UpstreamSpec{
					DiscoveryMetadata: &gloov1.DiscoveryMetadata{
						Labels: map[string]string{
							"app":     "httpbin",
							"service": "httpbin",
						},
					},
					UpstreamType: &gloov1.UpstreamSpec_Kube{
						Kube: &kubernetes.UpstreamSpec{
							Selector: map[string]string{
								"app": "httpbin",
							},
							ServiceNamespace: "httpbin",
							ServiceName:      "httpbin",
							ServicePort:      uint32(8000),
						},
					},
				},
			},
			&v1.VirtualService{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "httpbin-route",
					Namespace: "httpbin",
				},
				Spec: v1.VirtualServiceSpec{
					VirtualHost: &v1.VirtualHost{
						Domains: []string{"httpbin.example.com"},
						Routes: []*v1.Route{{
							Matchers: []*matchers.Matcher{
								{
									PathSpecifier: &matchers.Matcher_Prefix{
										Prefix: "/",
									},
								},
							},
							Action: &v1.Route_RouteAction{
								RouteAction: &gloov1.RouteAction{
									Destination: &gloov1.RouteAction_Single{
										Single: &gloov1.Destination{
											DestinationType: &gloov1.Destination_Upstream{
												Upstream: &gloocore.ResourceRef{
													Name:      "httpbin-htppbin-8000",
													Namespace: gloodefaults.GlooSystem,
												},
											},
										},
									},
								},
							},
						}},
					},
				},
			},
		}
		testcases.TestGatewayIngress(
			ctx,
			runner,
			testHelper,
			&snapshot.GatewayInfo{
				Name: "gateway-proxy",
				Port: 8080,
			},
			inputs,
			customSetupAssertions)
	})
})
