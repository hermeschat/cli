package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

//Get is a proxy to C().Get
var Get = C().Get

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
	"host": "", "port": "",
}
var config *ConfigMap

//C gets the global config object
func C() ConfigMap {
	return *config
}

/*FromEnv gets a map from env key to default value,
tries to get key from env if not found uses default value present as value of map */
func FromEnv(kd map[string]string) error {
	err := godotenv.Load()
	if err != nil {
		return errors.Wrap(err, "error in loading env")
	}
	if kd == nil {
		kd = envList
	}
	c := &ConfigMap{}
	for k, d := range kd {
		v := os.Getenv(strings.ToUpper(k))
		if v == "" {
			v = d
		}
		c.Set(strings.ToLower(k), v)
	}
	config = c
	return nil
}
