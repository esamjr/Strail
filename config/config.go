package config

import "github.com/tkanos/gonfig"

// Configuration
type Conf struct {
	DB_USER     string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

// setuo database
func GetConfig() Conf {
	config := Conf{}
	gonfig.GetConf("config/prod.json", &config) //Option *dev.json OR *prod.json
	return config
}
