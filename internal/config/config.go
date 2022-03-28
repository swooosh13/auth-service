package config

import (
	"fmt"
	"os"
	"path/filepath"
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
	SecretKey string `yaml:"secret"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		var configName string
		configName = os.Getenv("NODE_ENV")
		if configName == "" {
			configName = "local"
		}

		viper.SetConfigName(configName)
		viper.SetConfigType("yaml")

		dirPath, err := filepath.Abs("./configs")
		if err != nil {
			logger.Fatal(fmt.Sprintf("fatal error config dir: %s \n", err))
		}
		viper.AddConfigPath(dirPath)
		err = viper.ReadInConfig()
		if err != nil {
			logger.Fatal(fmt.Sprintf("fatal error config file: %s \n", err))
		}

		instance = &Config{}

		err = viper.Unmarshal(instance)
		if err != nil {
			logger.Fatal(fmt.Sprintf("fatal parse config: %s \n", err))
		}
	})

	return instance
}
