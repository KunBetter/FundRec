package config

import (
	"fmt"
	"github.com/KunBetter/FundRec/env"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Mysql Mysql `yaml:"mysql"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func LoadConfig() *Config {
	conf := &Config{}

	curEnv := env.GetCurEnv()
	configPath := "Config_Test.yaml"
	if curEnv == env.Prod {
		configPath = "Config_Prod.yaml"
	}

	if file, err := os.Open("config/" + configPath); err != nil {
		fmt.Print(err)
		return nil
	} else {
		err = yaml.NewDecoder(file).Decode(conf)
		if nil != err {
			fmt.Print(err)
		}
	}

	return conf
}
