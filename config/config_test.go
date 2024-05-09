package config

import (
	cm "code-kata/config/configManager"
	"testing"

	"github.com/stretchr/testify/suite"
)

type configTestSuite struct {
	suite.Suite
	*Config
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(configTestSuite))
}

func (suite *configTestSuite) SetupTest() {
	suite.Config = NewConfig()
	err := cm.NewConfigurationManager().Load("config_test.json", "./", suite.Config)
	if err != nil {
		suite.Fail("Could not load the config. ErrorL", err.Error())
	}
}

func (suite *configTestSuite) TestConfigValues() {
	suite.Equal("code-kata", suite.AppDetails.AppName)
	suite.Equal("1.0.0", suite.AppDetails.Version)
	suite.Equal(30, suite.HttpClientSettings.TimeOutInSecs)
	suite.Equal("https://jsonplaceholder.typicode.com/todos", suite.TodoApiDetails.Url)
}
