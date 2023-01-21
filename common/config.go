package common

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	Env            string             `mapstructure:"env"`
	Port           int                `mapstructure:"port"`
	Database       DatabaseCredential `mapstructure:"database"`
	ExternalClient ExternalClient     `mapstructure:"externalClient"`
}

func (c Configuration) GetPort() string {
	return fmt.Sprintf(":%d", c.Port)
}

type DatabaseCredential struct {
	Dialect  string `mapstructure:"dialect"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

type Veryfi struct {
	Url      string `mapstructure:"url"`
	ClientId string `mapstructure:"clientId"`
	Username string `mapstructure:"username"`
	ApiKey   string `mapstructure:"apiKey"`
}

type ExternalClient struct {
	Veryfi Veryfi `mapstructure:"veryfi"`
}

func ReadConfig(path string) (*Configuration, error) {
	var cfg Configuration

	logrus.Infof("reading %s cofiguration ...", path)

	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("failed to read config: %v", err)
		return nil, ErrConfigNotFound
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		logrus.Errorf("failed to marshal config: %v", err)
		return nil, ErrFailedMarshalingConfig
	}

	logrus.Infof("success reading %s cofiguration", path)
	return &cfg, nil
}
