package configuration

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"encoding/json"
)

type FileLoader struct {
	Path         string
	Unmarshaller UnmarshalFunc
}

type UnmarshalFunc func(data []byte, target interface{}) error

func NewFileLoader(path string, unmarshaller UnmarshalFunc) *FileLoader {
	return &FileLoader{path, unmarshaller}
}

func (k *FileLoader) Load() (map[string]interface{}, error) {
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

func NewYAMLFileLoader(path string) *FileLoader {
	return &FileLoader{path, yaml.Unmarshal}
}

func NewJSONFileLoader(path string) *FileLoader {
	return &FileLoader{path, json.Unmarshal}
}
