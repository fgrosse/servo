package configuration

type MemoryLoader struct {
	Data map[string]interface{}
}

func NewMemoryLoader() *MemoryLoader {
	return &MemoryLoader{map[string]interface{}{}}
}

func NewMemoryLoaderFrom(data map[string]interface{}) *MemoryLoader {
	l := NewMemoryLoader()
	l.SetAll(data)
	return l
}

func (k *MemoryLoader) Set(key string, value interface{}) {
	k.Data[key] = value
}

func (k *MemoryLoader) SetAll(data map[string]interface{}) {
	k.Data = data
}

func (k *MemoryLoader) Delete(key string) {
	delete(k.Data, key)
}

func (k *MemoryLoader) Load() (map[string]interface{}, error) {
	return k.Data, nil
}
