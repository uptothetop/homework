package blobfs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

const filename = "storage.json"

var storage = make(map[string][]byte)

func Get(key string) ([]byte, error) {
	loadStorage()
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
	saveStorage()
	return err
}

func Delete(key string) error {
	loadStorage()
	var err error
	defer func() {
		if result := recover(); result != nil {
			err = errors.New("error storing value")
		}
	}()

	delete(storage, key)
	saveStorage()
	return err
}

// Loads storage from the FS
func loadStorage() error {
	var err error

	jsonFile, err := os.Open("users.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &storage)

	return err
}

// Saves storage to the FS
func saveStorage() error {
	var err error

	jsonData, err := json.Marshal(storage)
	jsonFile, err := os.Create(filename)

	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()

	return err
}
