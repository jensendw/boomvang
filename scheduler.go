package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/jensendw/boomvang/collectors"
	//"github.com/jensendw/boomvang/scalers"
	"github.com/jensendw/boomvang/config"
	reflections "gopkg.in/oleiade/reflections.v1"
)

func collectorsScheduler() {
	if isConfigValueSet("RabbitMQUrl", myConfig) {
		rabbitMQConfig := collectors.RabbitMQClient{
			URL:      myConfig.RabbitMQUrl,
			Username: myConfig.RabbitMQUsername,
			Password: myConfig.RabbitMQPassword,
		}
		gocron.Every(myConfig.CollectionInterval).Seconds().Do(collectors.RunRabbitMQCollector, rabbitMQConfig)
	}

	if isConfigValueSet("DatadogAPIKey", myConfig) {
		datadogConfig := collectors.DatadogClient{
			URL:    myConfig.DatadogAPIEndpoint,
			APIKey: myConfig.DatadogAPIKey,
			AppKey: myConfig.DatadogAppKey,
		}
		gocron.Every(myConfig.CollectionInterval).Seconds().Do(collectors.RunDDCollector, datadogConfig)
	}

}

func isConfigValueSet(configKey string, theConfig *config.Config) bool {
	value, err := reflections.GetField(theConfig, configKey)
	if err != nil {
		return false
	}
	if value != "" {
		return true
	}
	return false
}

func RunSchedulers() {
	//providersScheduler()
	collectorsScheduler()
	<-gocron.Start()
}
