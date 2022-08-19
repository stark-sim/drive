package config

import (
	"drive/tools"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	DBConfig `mapstructure:"db"`

	Code `mapstructure:"code"`

	APIConfig `mapstructure:"api"`

	MinioConfig `mapstructure:"minio"`
}

type Code struct {
	Invite string
}

type APIConfig struct {
	Port int
}

type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type MinioConfig struct {
	Host      string
	Port      int
	AccessKey string
	SecretKey string
}

func Init() error {
	configPath := tools.GetDeployPath("/config.yaml")
	logrus.Printf("====> config path: %v", configPath)

	viper.SetConfigFile(configPath)
	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		logrus.Printf("config file has changed...")
		err := viper.Unmarshal(&Conf)
		if err != nil {
			logrus.Errorf("failed at unmarshal config file after change, err: %v", err)
			return
		}
	})

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("failed at ReadInConfig, err: %v", err)
	}

	if err = viper.Unmarshal(&Conf); err != nil {
		logrus.Errorf("failed at unmarshal config file, err: %v", err)
	}

	//viper.AutomaticEnv()

	return nil
}
