package service

import (
	"github.com/caarlos0/env"
)

// Config retrieves config from env vars that are specific to the
// server.
type Config struct {
	Addr    string `env:"ADDR" envDefault:":8080"`
	Logging bool   `env:"LOGGING" envDefault:"true"`
}

// New instantiates a new Config and attempts to parse parameters from
// corresponding environment variables.
func New() Config {
	mpc := Config{}
	env.Parse(&mpc)
	return mpc
}
