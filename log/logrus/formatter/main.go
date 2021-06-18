package main

import "github.com/sirupsen/logrus"

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: true,
	}
	logger.Infof("this is a message #1")
	logger.Formatter = &logrus.JSONFormatter{
		PrettyPrint: true,
	}
	logger.Infof("this is a message #2")
}
