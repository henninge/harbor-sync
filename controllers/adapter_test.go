package controllers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var _ = Describe("Adapter", func() {

	BeforeEach(func() {
		ensureHarborSyncConfig(k8sClient, "bar")
		ensureHarborSyncConfig(k8sClient, "baz")
		ensureHarborSyncConfig(k8sClient, "foo")
	})

	AfterEach(func() {
		deleteHarborSyncConfig(k8sClient, "bar")
		deleteHarborSyncConfig(k8sClient, "baz")
		deleteHarborSyncConfig(k8sClient, "foo")
	})

	It("should send events", func(done Done) {
		var ad Adapter
		input1 := make(chan struct{})
		input := []<-chan struct{}{input1}

		log := zap.Logger(false)

		ad = NewAdapter(k8sClient, log, input)
		out := ad.Run()

		input1 <- struct{}{}

		evt1 := <-out
		evt2 := <-out
		evt3 := <-out

		Expect(evt1.Meta.(*metav1.ObjectMeta).Name).To(Equal("bar"))
		Expect(evt2.Meta.(*metav1.ObjectMeta).Name).To(Equal("baz"))
		Expect(evt3.Meta.(*metav1.ObjectMeta).Name).To(Equal("foo"))
		close(done)
	})
})
