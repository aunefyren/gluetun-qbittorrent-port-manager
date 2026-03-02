package modules

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger(configFile ConfigStruct) {
	Log = logrus.New()

	// Define log file
	logFile, err := os.OpenFile("config/gluetun-qbittorrent-port-manager.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatalf("failed to load log file: %v", err)
	}

	// Set a plain text format with old-style timestamp
	Log.SetFormatter(&logrus.JSONFormatter{})

	// Output to both stdout and log file
	mw := io.MultiWriter(os.Stdout, logFile)
	Log.SetOutput(mw)

	// Set log level
	level, err := logrus.ParseLevel(configFile.LogLevel)
	if err != nil {
		logrus.Error("failed to load log file: %v", err)
		level = logrus.InfoLevel
	}

	Log.SetLevel(level)

	Log.Info("log level set to: " + level.String())
}
