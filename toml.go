package config

import (
	"github.com/BurntSushi/toml"
)

// parseToml parses a toml file and decodes it into the
// provided value, which must be passed as a pointer to
// some type that has already been allocated
func parseToml(fileName string, value interface{}) error {
	_, err := toml.DecodeFile(fileName, value)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}
