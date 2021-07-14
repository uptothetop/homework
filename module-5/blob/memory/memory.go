package blolbmemory

import "errors"

var storage = make(map[string][]byte)

func Get(key string) ([]byte, error) {
	val, err := storage[key]
	if err {
		return nil, errors.New("error getting data from the storage")
	}
	return val, nil
}

func Put(key string, value []byte) error {
	var err error
	defer func() {
		if result := recover(); result != nil {
			err = errors.New("error storing value")
		}
	}()

	storage[key] = value
	return err
}

func Delete(key string) error {
	var err error
	defer func() {
		if result := recover(); result != nil {
			err = errors.New("error storing value")
		}
	}()

	delete(storage, key)
	return err
}
