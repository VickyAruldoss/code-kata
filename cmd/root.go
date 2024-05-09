package cmd

import (
	"code-kata/arg"
	"code-kata/client"
	configuration "code-kata/config"
	cm "code-kata/config/configManager"
	"code-kata/todo/service"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "code-kata",
	Short: "command line tool for TODO's",
	Long:  "this is a cli which helps to print the todo's",
	Run: func(cmd *cobra.Command, args []string) {
		config := LoadConfiguration()
		baseHttpClient := &http.Client{Timeout: time.Duration(config.HttpClientSettings.TimeOutInSecs) * time.Second}
		httpHandler := client.NewHttpRequestHandler(config, baseHttpClient)
		toDoService := service.NewTodoService(httpHandler, config)
		todos, _ := toDoService.Get(rootCmdArgs.Limit, rootCmdArgs.Even)
		toDoService.Print(todos)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

var rootCmdArgs = arg.Args{
	Even:  true,
	Limit: 20,
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&rootCmdArgs.Limit, "limit", "l", 20, "The number of TODO's to limit to")
	rootCmd.PersistentFlags().BoolVarP(&rootCmdArgs.Even, "even", "e", true, "Whether to filter only even numbered TODO's or not")
}

func LoadConfiguration() configuration.Config {
	configuration := configuration.NewConfig()
	err := cm.NewConfigurationManager().Load("config.json", "./config", &configuration)
	if err != nil {
		log.Errorf("Error while loading configuration, error: %#v", err)
		os.Exit(1)
	}
	return *configuration
}
