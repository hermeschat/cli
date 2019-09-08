package main

import (
	"hermescli/cmd"

	"github.com/amirrezaask/config"
)

var envList = []string{
	"host", "port", "name",
}

func main() {
	config.Get("")
	cmd.Execute()
}
