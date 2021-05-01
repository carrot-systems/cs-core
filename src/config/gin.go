package config

import (
	env "github.com/carrot-systems/csl-env"
)

func GetWebPort() int {
	return env.RequireEnvInt("WEB_PORT")
}
