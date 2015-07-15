package servo

type ConfigurationFlattener struct{}

func NewConfigurationFlattener() *ConfigurationFlattener {
	return &ConfigurationFlattener{}
}

func (f *ConfigurationFlattener) Flatten(data map[string]interface{}) map[string]interface{} {
	m := map[string]interface{}{}
	f.flatten("", data, m)
	return m
}

func (f *ConfigurationFlattener) flatten(key string, value interface{}, m map[string]interface{}) {
	switch x := value.(type) {
	case map[string]interface{}:
		for childKey, childValue := range x {
			newKey := childKey
			if key != "" {
				newKey = key + "." + childKey
			}
			f.flatten(newKey, childValue, m)
		}
	default:
		m[key] = value
	}
}
