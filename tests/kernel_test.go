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

	Describe("Registering bundles", func() {
		It("It should register all types of the bundle", func() {
			kernel.Register(new(testAPI.TestBundle))
			Expect(kernel.TypeRegistry).To(HaveKey("test_bundle.my_type"))
		})

		PIt("It load the bundle configuration", func() {

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
			Expect(kernel.Run()).To(Succeed())
		})

		It("It should flatten all configuration parameters", func() {
			config.SetAll(map[string]interface{}{
				"nested": map[interface{}]interface{}{
					"foo": map[interface{}]interface{}{
						"value": "foo",
					},
					"bar": map[interface{}]interface{}{
						"value": "bar",
					},
				},
			})

			kernel.RegisterType("kernel.server", testAPI.NewServerMockWithParams, "%nested.foo.value%", "%nested.bar.value%")
			Expect(kernel.Run()).To(Succeed())
		})

		It("It should validate the container", func() {
			kernel.RegisterType("some.service", testAPI.NewRecursiveService, "@other.service")
			kernel.RegisterType("other.service", testAPI.NewRecursiveService, "@some.service")
			Expect(kernel.Run()).NotTo(Succeed(), "should return an error because we have a dependency cycle in the container")
		})

		It("It should use the kernel.server type to run the server", func() {
			Expect(kernel.Run()).To(Succeed())
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
