package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type QueueConifg struct {
	Handlers []string
}

type Configuration struct {
	Queues map[string]QueueConifg //Definition of queues
}

func Bootstrap() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	configuration := Configuration{
		Queues: make(map[string]QueueConifg, 0),
	}
}
