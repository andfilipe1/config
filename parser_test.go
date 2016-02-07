package config

import (
	"bytes"
	"encoding/json"
	"github.com/BurntSushi/toml"
	. "gopkg.in/check.v1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type ParserSuite struct{}

var _ = Suite(&ParserSuite{})

func (s *ParserSuite) TestParse(c *C) {
	type Nested struct {
		Parent string
	}

	type TestConfig struct {
		FileName string
		Count    int
		Child    Nested
	}

	config := TestConfig{
		FileName: "test.me",
		Count:    10,
		Child: Nested{
			Parent: "parent",
		},
	}

	exts := []string{"json", "toml", "yaml"}

	for _, ext := range exts {
		ext := ext
		func() {
			// we only need the temp filename
			file, err := ioutil.TempFile("", "test")
			c.Assert(err, IsNil)
			file.Close()

			var data []byte
			switch ext {
			case "json":
				data, err = json.Marshal(config)
				c.Assert(err, IsNil)
			case "toml":
				buf := bytes.NewBuffer(nil)
				err = toml.NewEncoder(buf).Encode(config)
				c.Assert(err, IsNil)
				data = buf.Bytes()
			case "yaml":
				data, err = yaml.Marshal(config)
				c.Assert(err, IsNil)
			}

			newFileName := file.Name() + "." + ext
			newFile, err := os.OpenFile(newFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
			c.Assert(err, IsNil)

			defer func() {
				_ = newFile.Close()
			}()

			n, err := newFile.Write(data)
			c.Assert(n, Equals, len(data))
			c.Assert(err, IsNil)

			// let's confirm that we can parse the data again
			config2 := TestConfig{}
			err = Parse(newFileName, &config2)
			c.Assert(err, IsNil)
			c.Assert(config2, DeepEquals, config)

			// shouldn't be able to parse other extensions
			err = Parse(file.Name(), &config2)
			c.Assert(err, NotNil)
			c.Assert(strings.HasPrefix(err.Error(), "unknown file format"), Equals, true)
		}()
	}
}
