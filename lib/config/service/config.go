package service

import (
	"github.com/caarlos0/env"
)

// Config retrieves config from env vars that are specific to the
// server.
type Config struct {
	Addr          string `env:"ADDR" envDefault:":8001"`
	Logging       bool   `env:"LOGGING" envDefault:"true"`
	StorageDriver string `env:"STORAGE_DRIVER" envDefault:"fs"`
}

// New instantiates a new Config and attempts to parse parameters from
// corresponding environment variables.
func New() Config {
	mpc := Config{}
	env.Parse(&mpc)
	return mpc
}
