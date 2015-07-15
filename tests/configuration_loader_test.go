package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"github.com/fgrosse/goku/vendor/src/gopkg.in/yaml.v2"
	"github.com/fgrosse/servo"
)

var _ = Describe("Build in configuration loaders", func() {

	Describe("ConfigurationFileLoader", func() {
		var config *servo.ConfigurationFileLoader

		BeforeEach(func() {
			config = servo.NewConfigurationFileLoader("fixtures/config.yml", yaml.Unmarshal)
		})

		Describe("Load", func() {
			It("return any errors that occur when opening the file", func() {
				config.Path = "/this/path/does/not/exist"
				_, err := config.Load()
				Expect(err).To(HaveOccurred())
			})

			It("return any errors that occur when unmarshalling the file", func() {
				config.Unmarshaller = func(data []byte, target interface{}) error {
					return fmt.Errorf("OH NO!")
				}

				_, err := config.Load()
				Expect(err).To(MatchError("error while parsing configuration file: OH NO!"))
			})

			It("should load the configuration using the provided file path and unmarshaller function", func() {
				data, err := config.Load()
				Expect(err).NotTo(HaveOccurred())
				Expect(data).To(HaveKeyWithValue("foo", "value1"))
				Expect(data).To(HaveKeyWithValue("bar", "value2"))
			})
		})
	})

	Describe("ConfigurationFlattener", func() {
		It("It should flatten all configuration parameters", func() {
			data := map[string]interface{}{
				"nested": map[string]interface{}{
					"foo": map[string]interface{}{
						"value": "nested value 1",
					},
					"bar": map[string]interface{}{
						"value": "nested value 2",
					},
				},
			}

			flattener := servo.NewConfigurationFlattener()
			data = flattener.Flatten(data)

			Expect(data).To(HaveKeyWithValue("nested.foo.value", "nested value 1"))
			Expect(data).To(HaveKeyWithValue("nested.bar.value", "nested value 2"))
		})
	})
})
