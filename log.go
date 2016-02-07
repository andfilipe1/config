package config

import (
	"github.com/op/go-logging"
)

const (
	module = "config"
)

var log *logging.Logger

func init() {
	log = logging.MustGetLogger(module)
}
