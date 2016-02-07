package config

import (
	"encoding/json"
	"os"
)

// parseJson parses a json file and decodes it into provided
// value, which must be passed as a pointer to
// some type that has already been allocated
func parseJson(fileName string, value interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(value)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}
