package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/mes1234/syncbrok/internal/space"
	"github.com/spf13/viper"
)

type HandlerConfig struct {
	url    string `mapstructure:"url"`
	queues string `mapstructure:"queues"`
}

type QueueConifg struct {
	Handlers []string
}

type Configuration struct {
	Queues map[string]QueueConifg //Definition of queues
}

func Bootstrap(
	wg *sync.WaitGroup,
	newMsgCh chan<- space.Messages,
	newSubscribersCh chan<- space.Subscribers,
	newQueueCh chan<- space.Queues) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	configuration := Configuration{
		Queues: make(map[string]QueueConifg),
	}
	configuration.initQueues(newQueueCh)

	configuration.initHandlers(newSubscribersCh)
}

func (c *Configuration) initHandlers(newSubscribersCh chan<- space.Subscribers) {
	handlers := make(map[string]*HandlerConfig)
	viper.UnmarshalKey("Handlers", &handlers)
	for index, value := range handlers {
		log.Printf("%v,%v", index, value)
	}
}

func (c *Configuration) initQueues(newQueueCh chan<- space.Queues) {

	queues := viper.GetStringSlice("Queues")

	for _, value := range queues {
		newQueue := QueueConifg{}
		c.Queues[value] = newQueue
		newQueueCh <- space.Queues{
			QName: value,
		}
	}

}
