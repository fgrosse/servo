package routing

type Route struct {
	EndpointTypeID string `yaml:"-"`
	Path           string `yaml:"path"`
}
