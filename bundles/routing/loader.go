package routing

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Loader struct {}

func (l *Loader) Load(filename string) (map[string]*Route, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error while loading routing file: %s", err)
	}

	routes := map[string]*Route{}
	err = yaml.Unmarshal(data, &routes)
	if err != nil {
		return nil, fmt.Errorf("error while parsing routing file: %s", err)
	}

	for name, r := range routes {
		if r.EndpointTypeID == "" {
			r.EndpointTypeID = name
		}
	}

	return routes, nil
}
