package setting

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Server struct {
	Port     		int `yaml:"port"`
	ReadTimeOut     int `yaml:"readTimeOut"`
	WriteTimeOut 	int `yaml:"writeTimeOut"`
}

type Config struct {
	Server *Server `yaml:"server"`
}


func InitConfig() Config{
	yamlFile, err := ioutil.ReadFile("./config/app.yml")
	if err != nil {
		log.Println(err.Error())
	}
	var _config *Config
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		log.Println(err.Error())
	}
	return *_config
}
