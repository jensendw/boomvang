package config

import (
	"github.com/jensendw/boomvang/logger"
	"github.com/spf13/viper"
	"strconv"
)

var Logger = *logger.Logger

//Config struct used to house configuration
type Config struct {
	RabbitMQUsername   string
	RabbitMQUrl        string
	RabbitMQPassword   string
	CollectionInterval uint64
	DatadogAPIKey      string
	DatadogAppKey      string
	DatadogAPIEndpoint string
}

//LoadConfig loads configuration when called
func LoadConfig() *Config {
	viper.SetEnvPrefix("BOOMVANG") // will be uppercased automatically
	viper.AutomaticEnv()

	// Set some defaults here
	viper.SetDefault("COLLECTION_INTERVAL", 60)
	viper.SetDefault("DATADOG_API_ENDPOINT", "https://app.datadoghq.com/api")

	//Convert the collection interval to uint64
	collectionInterval, err := strconv.ParseUint(viper.GetString("COLLECTION_INTERVAL"), 10, 32)
	if err != nil {
		Logger.Errorf("Issue converting collection interval %v", err)
	}

	config := Config{
		RabbitMQUrl:        viper.GetString("RABBITMQ_URL"),
		RabbitMQUsername:   viper.GetString("RABBITMQ_USERNAME"),
		RabbitMQPassword:   viper.GetString("RABBITMQ_PASSWORD"),
		DatadogAPIKey:      viper.GetString("DATADOG_API_KEY"),
		DatadogAppKey:      viper.GetString("DATADOG_APP_KEY"),
		DatadogAPIEndpoint: viper.GetString("DATADOG_API_ENDPOINT"),
		CollectionInterval: collectionInterval,
	}

	return &config
}
