package servo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fgrosse/servo"
	"github.com/fgrosse/servo/configuration"
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
			Expect(kernel.TypeRegistry).To(HaveKey("kernel.http.server"))
		})
	})

	Describe("Registering bundles", func() {
		It("It should register all types of the bundle", func() {
			kernel.Register(new(TestBundle))
			Expect(kernel.TypeRegistry).To(HaveKey("test_bundle.my_type"))
		})

		PIt("It load the bundle configuration", func() {

		})
	})

	Describe("Run", func() {
		var server *ServerMock

		BeforeEach(func() {
			server = new(ServerMock)
			kernel.InjectInstance("kernel.http.server", server)
			config.Set("servo.listen", "0.0.0.0:3000")
		})

		It("It should load the configuration", func() {
			kernel.RegisterType("kernel.http.server", NewServerMockWithParams, "%foo%", "%bar%")
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

			kernel.RegisterType("kernel.http.server", NewServerMockWithParams, "%nested.foo.value%", "%nested.bar.value%")
			Expect(kernel.Run()).To(Succeed())
		})

		It("It should validate the container", func() {
			kernel.RegisterType("some.service", NewRecursiveService, "@other.service")
			kernel.RegisterType("other.service", NewRecursiveService, "@some.service")
			Expect(kernel.Run()).NotTo(Succeed(), "should return an error because we have a dependency cycle in the container")
		})

		It("It should use the kernel.http.server type to run the server", func() {
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
