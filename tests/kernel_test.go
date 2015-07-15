package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fgrosse/servo"
	"github.com/fgrosse/servo/tests/testAPI"
)

var _ = Describe("Kernel", func() {
	var (
		config *servo.MemoryConfigurationLoader
		kernel *servo.Kernel
	)

	BeforeEach(func() {
		config = servo.NewMemoryConfigurationLoader()
		kernel = servo.NewKernel(config)
	})

	Describe("NewKernel", func() {
		It("It should initialize the goldi.TypeRegistry", func() {
			Expect(kernel.TypeRegistry).NotTo(BeNil())
		})

		It("It should register all internal types", func() {
			Expect(kernel.TypeRegistry).To(HaveKey("kernel.server"))
		})
	})

	Describe("Run", func() {
		var server *testAPI.ServerMock

		BeforeEach(func() {
			server = new(testAPI.ServerMock)
			kernel.InjectInstance("kernel.server", server)
		})

		It("It should load the configuration", func() {
			kernel.RegisterType("kernel.server", testAPI.NewServerMockWithParams, "%foo%", "%bar%")
			config.Set("foo", "foo")
			config.Set("bar", "bar")
			kernel.Run()
		})

		PIt("It should validate the container", func() {
			// TODO
		})

		It("It should use the kernel.server type to run the server", func() {
			kernel.Run()
			Expect(server.RunHasBeenCalled).To(BeTrue())
		})

		It("It should return errors from server.Run", func() {
			server.ReturnError = false
			result := kernel.Run()
			Expect(result).To(BeNil())

			server.ReturnError = true
			result = kernel.Run()
			Expect(result).NotTo(BeNil())
		})
	})
})
