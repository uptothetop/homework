package blolbmemory

// Exported to use it from FS blob
var Storage = make(map[string][]byte)

func Get(key string) ([]byte, error) {
	return Storage[key], nil // TODO: Return errors
}

func Put(key string, value []byte) error {
	Storage[key] = value
	return nil // TODO: Return errors
}

func Delete(key string) error {
	delete(Storage, key)
	return nil // TODO: Return errors
}
