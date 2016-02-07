package config

import (
	"github.com/BurntSushi/toml"
	. "gopkg.in/check.v1"
	"io/ioutil"
)

type TomlSuite struct{}

var _ = Suite(&TomlSuite{})

func (s *TomlSuite) TestParse(c *C) {
	type Nested struct {
		Parent string `toml:"parent"`
	}

	type TestConfig struct {
		FileName string `toml:"file_name"`
		Count    int    `toml:"count"`
		Child    Nested `toml:"child"`
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

	err = toml.NewEncoder(file).Encode(config)
	c.Assert(err, IsNil)

	// let's confirm that we can parse the data again
	config2 := TestConfig{}
	err = parseToml(file.Name(), &config2)
	c.Assert(err, IsNil)
	c.Assert(config2, DeepEquals, config)
}
