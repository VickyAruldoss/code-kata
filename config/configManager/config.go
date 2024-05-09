package configManager

import (
	"encoding/json"

	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type ConfigurationManager interface {
	Load(fileName string, path string, configuration interface{}) error
}

type configurationManager struct {
}

func NewConfigurationManager() ConfigurationManager {
	return configurationManager{}
}

func (cm configurationManager) Load(fileName string, path string, configuration interface{}) (err error) {
	jsonfile, err := os.Open(filepath.Join(path, fileName))

	if err != nil {
		log.Errorf("Error while opening config file : %v", err)
		return
	}
	defer jsonfile.Close()

	decoder := json.NewDecoder(jsonfile)
	err = decoder.Decode(configuration)
	if err != nil {
		log.Error("Error while decoding json, Error:", err)
		return
	}
	return
}
