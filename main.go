package main

import (
	"github.com/BalanSnack/BACKEND/conf"
	"github.com/BalanSnack/BACKEND/handler"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	conf.Setup()
	// logrus setting
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)

	r := handler.NewGinEngine()

	r.Run("localhost:5000")
}
