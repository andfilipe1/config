package config

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"io/ioutil"
)

type JsonSuite struct{}

var _ = Suite(&JsonSuite{})

func (s *JsonSuite) TestParse(c *C) {
	type Nested struct {
		Parent string `json:"parent"`
	}

	type TestConfig struct {
		FileName string `json:"file_name"`
		Count    int    `json:"count"`
		Child    Nested `json:"child"`
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

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&config)
	c.Assert(err, IsNil)

	// let's confirm that we can parse the data again
	config2 := TestConfig{}
	err = parseJson(file.Name(), &config2)
	c.Assert(err, IsNil)
	c.Assert(config2, DeepEquals, config)
}
