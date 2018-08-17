package collectors

import (
	"github.com/jensendw/boomvang/logger"
)

var Logger = *logger.Logger

type MetricNames struct {
	MetricName      string
	MetricCollector string
}

// func writeMetricNames() error {
// 	//add the database stuff here
//
// }
