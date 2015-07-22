package configuration

import "reflect"

type Flattener struct {}

func (f *Flattener) Flatten(config map[string]interface{}) map[string]interface{} {
	flattenedConfig := map[string]interface{}{}
	f.flatten("", config, flattenedConfig)

	return flattenedConfig
}

func (f *Flattener) flatten(key string, value interface{}, m map[string]interface{}) {
	reflectedValue := reflect.ValueOf(value)
	switch reflectedValue.Kind() {
	case reflect.Map:
		for _, childKey := range reflectedValue.MapKeys() {
			childValue := reflectedValue.MapIndex(childKey).Interface()

			newKey := childKey.Interface().(string)
			if key != "" {
				newKey = key + "." + newKey
			}
			f.flatten(newKey, childValue, m)
		}
	default:
		m[key] = value
	}
}
