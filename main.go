package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("hello")

	//secret leak
	aws_secret := "AKIAIMNOJVGFDXXXE4OA"
	logrus.Info(aws_secret)

}
