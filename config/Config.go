package config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
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

func ReadYamlConfig(path string) (*Config, error) {
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	return conf, nil
}

func ConfigTest() {
	conf, err := ReadYamlConfig("D:/test_yaml/test.yaml")
	if err != nil {
		log.Fatal(err)
	}

	byts, err := json.Marshal(conf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byts))
}
