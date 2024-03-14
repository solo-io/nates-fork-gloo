package testcases

import (
	"context"
	"fmt"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/test/kube2e/helper"

	"github.com/solo-io/gloo/test/snapshot"
	"github.com/solo-io/skv2/codegen/util"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var TestGatewayIngress = func(
	ctx context.Context,
	runner snapshot.TestRunner,
	testHelper *helper.SoloTestHelper,
	gatewayInfo *snapshot.GatewayInfo,
	inputs []client.Object,
	customSetupAssertions func(),
) {
	It("should translate a gateway with basic routing", func() {
		dir := util.MustGetThisDir()
		runner.ResultsByGateway = map[types.NamespacedName]snapshot.ExpectedTestResult{
			{
				Namespace: "default",
				Name:      "example-gateway-http",
			}: {
				Proxy: dir + "/outputs/http-routing-proxy.yaml",
				// Reports:     nil,
			},
		}

		err := runner.Run(ctx, inputs)
		Expect(err).NotTo(HaveOccurred())

		// check setup assertions
		customSetupAssertions()

		testHelper.CurlEventuallyShouldRespond(helper.CurlOpts{
			Protocol:          "http",
			Path:              "/headers",
			Method:            "GET",
			Host:              "httpbin.example.com",
			Service:           gatewayInfo.Name,
			Port:              gatewayInfo.Port,
			ConnectionTimeout: 10,
			Verbose:           true,
			WithoutStats:      true,
			ReturnHeaders:     true,
		}, fmt.Sprintf("HTTP/1.1 %d", http.StatusOK), 1, time.Minute*1)
	})
}
