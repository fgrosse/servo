package servo

type ConfigurationLoader interface {
	Load() (map[string]interface{}, error)
}
