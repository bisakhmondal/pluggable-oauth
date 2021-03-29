package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	Credentials struct {
		Google   Creds
		Facebook Creds
		Github   Creds
		Linkedin Creds
	}
}

type Creds struct {
	Id     string
	Secret string
}

func (c *Config) Parse() {
	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(c)
	if err != nil {
		panic(err)
	}
}
