package configuration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"github.com/fgrosse/servo/configuration"
)

var _ = Describe("FileLoader", func() {
	var config *configuration.FileLoader

	BeforeEach(func() {
		config = configuration.NewYAMLFileLoader("../_test_fixtures/config.yml")
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
