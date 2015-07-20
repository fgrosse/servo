package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fgrosse/servo"
	"github.com/fgrosse/servo/configuration"
	"github.com/fgrosse/servo/tests/testAPI"
)

var _ = Describe("Kernel", func() {
	var (
		config *configuration.MemoryLoader
		kernel *servo.Kernel
	)

	BeforeEach(func() {
		config = configuration.NewMemoryLoader()
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

		It("It should flatten all configuration parameters", func() {
			config.SetAll(map[string]interface{}{
				"nested": map[string]interface{}{
					"foo": map[string]interface{}{
						"value": "foo",
					},
					"bar": map[string]interface{}{
						"value": "bar",
					},
				},
			})

			kernel.RegisterType("kernel.server", testAPI.NewServerMockWithParams, "%nested.foo.value%", "%nested.bar.value%")
			kernel.Run()
		})

		It("It should validate the container", func() {
			kernel.RegisterType("some.service", testAPI.NewRecursiveSerice, "@other.service")
			kernel.RegisterType("other.service", testAPI.NewRecursiveSerice, "@some.service")
			Expect(func() { kernel.Run() }).To(Panic(), "should panic because we have a dependency cycle in the container")
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
