package routing

type Route struct {
	EndpointTypeID string `yaml:"endpoint"`
	Path           string `yaml:"path"`
}
