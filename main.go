package main

import (
	"hermescli/cmd"
	"hermescli/config"
)

var envList = map[string]string{
	"host": "localhost", "port": "9000",
}

func main() {
	config.Init(envList)
	cmd.Execute()
}
