package main

import (
	"github.com/jensendw/boomvang/config"
)

var myConfig = config.LoadConfig()

func main() {
	RunSchedulers()
}
