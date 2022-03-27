package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"github.com/swooosh13/quest-auth/pkg/logger"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"listen"`
	MongoDB struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Database   string `yaml:"database"`
		AuthDB     string `yaml:"authdb"`
		Collection string `yaml:"collection"`
	} `yaml:"mongodb"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("local")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./configs")

		err := viper.ReadInConfig()
		if err != nil {
			logger.Fatal(fmt.Sprint("fatal error config file: %w \n", err))
		}

		instance = &Config{}

		err = viper.Unmarshal(instance)
		if err != nil {
			logger.Fatal(fmt.Sprint("Fatal parse config: %w \n", err))
		}
	})

	return instance
}
