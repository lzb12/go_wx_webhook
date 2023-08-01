package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Yaml struct {

}

type Config struct {
	Server 	ServerConfig
	Redis 	CRedis
}
type CRedis struct {
	Host 		string  `yaml:"host"`
	Port 		int		`yaml:"port"`
	Password 	string 	`yaml:"password"`
	db 			int 	`yaml:"db"`
}

type ServerConfig struct {
	Host 	string 	`yaml:"host"`
	Port 	int		`yaml:"port"`
}


var Conf Config

func (y *Yaml) LoadToml()  {

	yaml := viper.New()
	yaml.SetConfigName("app")
	yaml.SetConfigType("yaml")
	yaml.AddConfigPath("config")
	if err := yaml.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}


	err := yaml.Unmarshal(&Conf)
	if err != nil {
		fmt.Println(err)
		return
	}




}