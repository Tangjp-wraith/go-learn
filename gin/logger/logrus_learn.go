package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	//logrus.SetLevel(logrus.ErrorLevel)
	//logrus.Error("error")
	//logrus.Warnln("warn")
	//logrus.Infof("info")
	//logrus.Debugf("debug")
	//logrus.Println("print")
	//fmt.Println(logrus.GetLevel())

	logrus.SetFormatter(&logrus.JSONFormatter{})
	log := logrus.WithField("app", "study").WithField("service", "fmy")
	log = log.WithFields(logrus.Fields{
		"user_id": "21",
		"ip":      "0.0.0.1",
	})

	log.Error("sb")
}
