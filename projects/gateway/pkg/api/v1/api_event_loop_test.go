package v1

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
)

var _ = Describe("ApiEventLoop", func() {
	var (
		namespace string
		emitter   ApiEmitter
		err       error
	)

	BeforeEach(func() {

		gatewayClientFactory := factory.NewResourceClientFactory(&factory.MemoryResourceClientOpts{
			Cache: memory.NewInMemoryResourceCache(),
		})
		gatewayClient, err := NewGatewayClient(gatewayClientFactory)
		Expect(err).NotTo(HaveOccurred())

		virtualServiceClientFactory := factory.NewResourceClientFactory(&factory.MemoryResourceClientOpts{
			Cache: memory.NewInMemoryResourceCache(),
		})
		virtualServiceClient, err := NewVirtualServiceClient(virtualServiceClientFactory)
		Expect(err).NotTo(HaveOccurred())

		emitter = NewApiEmitter(gatewayClient, virtualServiceClient)
	})
	It("runs sync function on a new snapshot", func() {
		_, err = emitter.Gateway().Write(NewGateway(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.VirtualService().Write(NewVirtualService(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		sync := &mockApiSyncer{}
		el := NewApiEventLoop(emitter, sync)
		_, err := el.Run([]string{namespace}, clients.WatchOpts{})
		Expect(err).NotTo(HaveOccurred())
		Eventually(func() bool { return sync.synced }, time.Second).Should(BeTrue())
	})
})

type mockApiSyncer struct {
	synced bool
}

func (s *mockApiSyncer) Sync(ctx context.Context, snap *ApiSnapshot) error {
	s.synced = true
	return nil
}
