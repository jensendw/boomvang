package main

import (
	"github.com/jensendw/boomvang/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIsConfigValueSet(t *testing.T) {
	os.Setenv("COLLECTION_INTERVAL", "10")
	myConfig := config.Config{
		RabbitMQUrl:   "https://rabbitmq:15672",
		DatadogAPIKey: "",
	}

	assert.True(t, isConfigValueSet("RabbitMQUrl", &myConfig))
	assert.False(t, isConfigValueSet("DatadogAPIKey", &myConfig))
	assert.False(t, isConfigValueSet("KeyThatDoesNotExist", &myConfig))

}
