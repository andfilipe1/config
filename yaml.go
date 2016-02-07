package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// parseYaml parses a yaml file and decodes it into the
// provided value, which must be passed as a pointer to
// some type that has already been allocated
func parseYaml(fileName string, value interface{}) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	err = yaml.Unmarshal(data, value)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}
