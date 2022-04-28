package main

import (
	_ "github.com/crewjam/saml"
	_ "github.com/prometheus/client_golang/api"
	_ "github.com/rehacktive/caffeine/database"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("hello")

	//secret leak
	aws_secret := "AKIAIMNOJVGFDXXXE4OA"
	logrus.Info(aws_secret)

	custom := "abcdefgh"
	logrus.Info(custom)
}
