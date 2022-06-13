package logger

import (
	"os"

	"github.com/MalukiMuthusi/cryptocurrencies/internal/utils"
	"github.com/sirupsen/logrus"
)

// Log is the main logger
var Log *logrus.Logger

func init() {
	Log = logrus.New()

	env := os.Getenv(utils.AppName)

	Log.SetReportCaller(true)

	if env == "dev" {
		Log.SetLevel(logrus.DebugLevel)
	}

}
