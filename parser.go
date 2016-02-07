package config

import (
	"fmt"
	"strings"
)

// Parse parses the file and decodes it into the pass value,
// which must be passed as a pointer to some type that
// has already been allocated.
// Depending on the file extension (.json, .toml, or .yaml), it will
// use the right parser
func Parse(fileName string, value interface{}) (err error) {
	lower := strings.ToLower(fileName)
	switch {
	case strings.HasSuffix(lower, ".json"):
		return parseJson(fileName, value)
	case strings.HasSuffix(lower, ".yaml"):
		return parseYaml(fileName, value)
	case strings.HasSuffix(lower, ".toml"):
		return parseToml(fileName, value)
	default:
		err = fmt.Errorf("unknown file format for " + fileName)
	}
	return err
}
