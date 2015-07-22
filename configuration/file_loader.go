package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type FileLoader struct {
	Path         string
	Unmarshaller UnmarshalFunc
}

type UnmarshalFunc func(data []byte, target interface{}) error

func NewFileLoader(path string, unmarshaller UnmarshalFunc) *FileLoader {
	return &FileLoader{path, unmarshaller}
}

func (l *FileLoader) Load() (map[string]interface{}, error) {
	err := l.determineAbsolutePath()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(l.Path)
	if err != nil {
		return nil, fmt.Errorf("error while loading configuration file: %s", err)
	}

	config := map[string]interface{}{}
	err = l.Unmarshaller(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error while parsing configuration file: %s", err)
	}

	return config, nil
}

func (l *FileLoader) determineAbsolutePath() error {
	if filepath.IsAbs(l.Path) {
		return nil
	}

	defer func() {
		l.Path, _ = filepath.Abs(l.Path)
	}()

	_, err := os.Stat(l.Path)
	if err == nil {
		return err
	}

	if os.IsNotExist(err) {
		absPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		absPath = absPath + "/" + l.Path

		_, err := os.Stat(absPath)
		if err == nil {
			l.Path = absPath
		}

		return nil
	}

	return fmt.Errorf("error while retrieving file stat: %s", err)
}

func NewYAMLFileLoader(path string) *FileLoader {
	return &FileLoader{path, yaml.Unmarshal}
}

func NewJSONFileLoader(path string) *FileLoader {
	return &FileLoader{path, json.Unmarshal}
}
