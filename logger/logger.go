package logger

import "github.com/op/go-logging"

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level} %{color:reset} %{message}`,
)

func init() {
	logging.SetFormatter(format)
}

//Logger used to send logs
var Logger = logging.MustGetLogger("boomvang")
