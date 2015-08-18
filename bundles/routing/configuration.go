package routing

type Route struct {
	Path           string `yaml:"path"`
	EndpointTypeID string `yaml:"endpoint"`
}
