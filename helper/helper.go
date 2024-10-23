package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	DBHost     string `json:"db"`
	DBPort     string `json:"port"`
	DBUsername string `json:"username"`
	DBPassword string `json:"password"`
	DBName     string `json:"database"`
}

var config *Config
var ConfigFile string

func ValidateConfig() bool {
	if ConfigFile == "" {
		return false
	}
	config := GetConfig()
	if config.DBHost == "" || config.DBPort == "" || config.DBUsername == "" || config.DBPassword == "" || config.DBName == "" {
		return false
	}
	return true
}

func GetConfig() *Config {
	//fmt.Println("ConfigFile", ConfigFile)
	if config == nil {
		config = NewConfig()
		if ConfigFile != "" {
			//fmt.Println(ConfigFile)
			config.Load(ConfigFile)
		}
	}
	return config
}

func NewConfig() *Config {
	c := Config{}
	return &c
}

func (c *Config) Load(configfile string) error {
	jsonFile, err := os.Open(configfile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = json.Unmarshal(byteValue, c)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
