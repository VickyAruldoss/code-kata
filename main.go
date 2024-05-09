package main

import (
	"code-kata/cmd"
	"code-kata/config"
	cm "code-kata/config/configManager"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	configuration := config.NewConfig()
	err := cm.NewConfigurationManager().Load("config.json", "./config", &configuration)
	if err != nil {
		log.Errorf("Error while loading configuration, error: %#v", err)
		os.Exit(1)
	}
	root := cmd.NewRoot(*configuration)
	root.Execute()
}
