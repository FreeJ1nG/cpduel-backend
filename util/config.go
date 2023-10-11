package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Env        string `mapstructure:"ENV"`
	ServerPort string `mapstructure:"SERVER_PORT"`

	Headless bool `mapstructure:"HEADLESS"`

	DBUser         string `mapstructure:"DB_USER"`
	DBUserPassword string `mapstructure:"DB_USER_PASSWORD"`

	DBHost string `mapstructure:"DB_HOST"`
	DBName string `mapstructure:"DB_NAME"`
	DBPort string `mapstructure:"DB_PORT"`

	DBDsn string
}

func SetConfig() (config Config, err error) {
	viper.AddConfigPath("../")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	config.setDBConfig()

	return
}

func (c *Config) setDBConfig() {
	var ssl string
	if c.Env == "local" {
		ssl = "sslmode=disable"
	} else {
		ssl = "sslmode=require"
	}

	DBDsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s %s",
		c.DBHost,
		c.DBUser,
		c.DBUserPassword,
		c.DBName,
		c.DBPort,
		ssl,
	)

	c.DBDsn = DBDsn
}
