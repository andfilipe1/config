package config

import (
	. "gopkg.in/check.v1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type YamlSuite struct{}

var _ = Suite(&YamlSuite{})

func (s *YamlSuite) TestParse(c *C) {
	type Nested struct {
		Parent string `yaml:"parent"`
	}

	type TestConfig struct {
		FileName string `yaml:"file_name"`
		Count    int    `yaml:"count"`
		Child    Nested `yaml:"child"`
	}

	config := TestConfig{
		FileName: "test.me",
		Count:    10,
		Child: Nested{
			Parent: "parent",
		},
	}

	file, err := ioutil.TempFile("", "test")
	c.Assert(err, IsNil)

	defer func() {
		_ = file.Close()
	}()

	data, err := yaml.Marshal(config)
	c.Assert(err, IsNil)
	c.Logf("content: %s\n", string(data))
	_, err = file.Write(data)
	c.Assert(err, IsNil)

	// let's confirm that we can parse the data again
	config2 := TestConfig{}
	err = parseYaml(file.Name(), &config2)
	c.Assert(err, IsNil)
	c.Assert(config2, DeepEquals, config)
}
