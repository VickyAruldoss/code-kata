package configManager

import (
	"code-kata/config"
	"testing"

	"github.com/stretchr/testify/suite"
)

type configManagerTestSuite struct {
	suite.Suite
	configLoader ConfigurationManager
}

func TestConfigurationManagerTestSuite(t *testing.T) {
	suite.Run(t, new(configManagerTestSuite))
}

func (suite *configManagerTestSuite) SetupTest() {
	suite.configLoader = NewConfigurationManager()
}

func (suite *configManagerTestSuite) TestLoad_Should_Return_Config_Data_When_File_Exists() {
	expectedConfig := &config.Config{
		AppDetails:         config.AppDetails{AppName: "code-kata", Version: "1.0.0"},
		HttpClientSettings: config.HttpClientSettings{TimeOutInSecs: 30},
		TodoApiDetails:     config.TodoApiDetails{Url: "https://jsonplaceholder.typicode.com/todos"},
	}
	configData := config.NewConfig()
	gotErr := suite.configLoader.Load("config_test.json", "../", configData)

	suite.Nil(gotErr)
	suite.NotNil(configData)
	suite.Equal(expectedConfig, configData)
}

func (suite *configManagerTestSuite) TestLoad_Should_Throw_Error_When_File_Doesnt_Exists() {

	configData := config.NewConfig()
	gotErr := suite.configLoader.Load("config_test.json", "../config", configData)

	suite.NotNil(gotErr)
}

func (suite *configManagerTestSuite) TestLoad_Should_Throw_Error_When_content_in_json_is_invalid() {

	configData := config.NewConfig()
	gotErr := suite.configLoader.Load("invalidContent.json", "../", configData)

	suite.NotNil(gotErr)
	suite.Equal("invalid character 's' after object key", gotErr.Error())
}
