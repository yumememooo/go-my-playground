package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Name string
	Port int

	File  string
	Level string
}

var C Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	if err := viper.Unmarshal(&C); err != nil {
		log.Printf("unable to decode into struct, %v", err)
	}

	log.Printf("read:%s", C.Name)
}
