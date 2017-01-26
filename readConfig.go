package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Database struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Name string `yaml:"name"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	}
	LDAP struct {
		Host        []string `yaml:"host"`
		ProxyBaseDN string   `yaml:"proxyBaseDN"`
		ProxyUserDN string   `yaml:"proxyUserDN"`
		ProxyPass   string   `yaml:"proxyPass"`
	}
}

func readConfig() Config {

	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Println(err)
	}

	c := Config{}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Println(err)
	}

	return c
}
