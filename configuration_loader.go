package servo

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"encoding/json"
)

type ConfigurationLoader interface {
	Load() (map[string]interface{}, error)
}

type MemoryConfigurationLoader struct {
	Data map[string]interface{}
}

func NewMemoryConfigurationLoader() *MemoryConfigurationLoader {
	return &MemoryConfigurationLoader{map[string]interface{}{}}
}

func NewMemoryConfigurationLoaderFrom(data map[string]interface{}) *MemoryConfigurationLoader {
	l := NewMemoryConfigurationLoader()
	l.SetAll(data)
	return l
}

func (k *MemoryConfigurationLoader) Set(key string, value interface{}) {
	k.Data[key] = value
}

func (k *MemoryConfigurationLoader) SetAll(data map[string]interface{}) {
	k.Data = data
}

func (k *MemoryConfigurationLoader) Delete(key string) {
	delete(k.Data, key)
}

func (k *MemoryConfigurationLoader) Load() (map[string]interface{}, error) {
	return k.Data, nil
}

type ConfigurationFileLoader struct {
	Path         string
	Unmarshaller UnmarshalFunc
}

type UnmarshalFunc func(data []byte, target interface{}) error

func NewYAMLFileLoader(path string) *ConfigurationFileLoader {
	return &ConfigurationFileLoader{path, yaml.Unmarshal}
}

func NewJSONFileLoader(path string) *ConfigurationFileLoader {
	return &ConfigurationFileLoader{path, json.Unmarshal}
}

func NewConfigurationFileLoader(path string, unmarshaller UnmarshalFunc) *ConfigurationFileLoader {
	return &ConfigurationFileLoader{path, unmarshaller}
}

func (k *ConfigurationFileLoader) Load() (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(k.Path)
	if err != nil {
		return nil, fmt.Errorf("error while loading configuration file: %s", err)
	}

	config := map[string]interface{}{}
	err = k.Unmarshaller(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error while parsing configuration file: %s", err)
	}

	return config, nil
}

