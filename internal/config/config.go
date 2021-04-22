package config

import (
	"fmt"
	"sync"

	"github.com/mes1234/syncbrok/internal/space"
	"github.com/spf13/viper"
)

type QueueConifg struct {
	Urls []string `yaml:"url"`
}

type Configuration struct {
	Queues           map[string]QueueConifg //Definition of queues
	newMsgCh         chan<- space.Message
	newSubscribersCh chan<- space.Subscriber
	newQueueCh       chan<- space.Queue
}

func Bootstrap(
	wg *sync.WaitGroup,
	newMsgCh chan<- space.Message,
	newSubscribersCh chan<- space.Subscriber,
	newQueueCh chan<- space.Queue) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	configuration := Configuration{
		Queues:           make(map[string]QueueConifg),
		newMsgCh:         newMsgCh,
		newQueueCh:       newQueueCh,
		newSubscribersCh: newSubscribersCh,
	}
	configuration.initQueues()
}

func (c *Configuration) initQueues() {
	viper.Unmarshal(&c)
	for queue, config := range c.Queues {
		c.newQueueCh <- space.Queue{QName: queue}
		for _, value := range config.Urls {
			c.newSubscribersCh <- space.Subscriber{
				QName:    queue,
				Endpoint: value,
			}
		}

	}
}
