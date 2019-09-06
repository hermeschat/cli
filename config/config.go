package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

//Get is a proxy to C().Get
var Get func(key string) string

//ConfigMap represnets config
type ConfigMap map[string]string

func (c ConfigMap) Get(key string) string {
	val, exists := (c)[key]
	if !exists {
		return ""
	}
	return val
}

func (c *ConfigMap) Set(key, value string) {
	(*c)[key] = value
}

//envlist shows whats envs are available
var envList = map[string]string{
	"host": "", "port": "", "name": "niima",
}
var config *ConfigMap

//C gets the global config object
func C() ConfigMap {
	return *config
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprint(os.Stderr, errors.Wrap(err, "error in loading env").Error())
	}
	c := &ConfigMap{}
	for k, d := range envList {
		v := os.Getenv(strings.ToUpper(k))
		if v == "" {
			v = d
		}
		c.Set(strings.ToLower(k), v)
	}
	config = c
	Get = config.Get
}
