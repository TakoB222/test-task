package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	defaultHttpPort      = "8000"
	defaultHttpRWTimeout = 10 * time.Second

	defaultConfigPath = "../config/config.yml"
)

type (
	Config struct {
		Http HttpServer
	}

	HttpServer struct {
		Host         string        `mapstructure:"host"`
		Port         string        `mapstructure:"port"`
		ReadTimeout  time.Duration `mapstructure:"readTimeout"`
		WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	}
)

func Init(path string) (*Config, error) {
	if path == "" {
		path = defaultConfigPath
	}
	populateDefaults()

	if err := parseConfigFile(path); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func populateDefaults() {
	viper.SetDefault("http.host", "localhost:"+defaultHttpPort)
	viper.SetDefault("http.readTimeout", defaultHttpRWTimeout)
	viper.SetDefault("http.writeTimeout", defaultHttpRWTimeout)
}

func parseConfigFile(filePath string) error {
	viper.AddConfigPath("config") // folder
	viper.SetConfigName("config") // config file name

	return viper.ReadInConfig()
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.Http); err != nil {
		return err
	}

	return nil
}
