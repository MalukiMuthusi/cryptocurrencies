package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Log is the main logger
var Log *logrus.Logger

func init() {
	Log = logrus.New()

	env := os.Getenv("CRYPTOCURRENCIES")

	Log.SetReportCaller(true)

	if env == "dev" {
		Log.SetLevel(logrus.DebugLevel)
	}

}
