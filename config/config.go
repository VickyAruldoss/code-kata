package config

func NewConfig() *Config {
	return &Config{}
}

type Config struct {
	AppDetails         AppDetails         `json:"app_details" validate:"required"`
	HttpClientSettings HttpClientSettings `json:"http_client_settings" validate:"required"`
	TodoApiDetails     TodoApiDetails     `json:"todo_api_details" validate:"required"`
}

type AppDetails struct {
	AppName string `json:"appName" validate:"required"`
	Version string `json:"version" validate:"required"`
}

type HttpClientSettings struct {
	TimeOutInSecs int `json:"timeOutInSecs" validate:"required"`
}

type TodoApiDetails struct {
	Url string `json:"url" validate:"required"`
}
