package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME      string
	DB_PASSWORD      string
	DB_HOST          string
	DB_PORT          string
	DB_NAME          string
	SMTP_HOST        string
	SMTP_PORT        string
	SMTP_SENDER_NAME string
	SMTP_USERNAME    string
	SMTP_PASSWORD    string
}

func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("config/config.json", &conf)
	return conf
}
