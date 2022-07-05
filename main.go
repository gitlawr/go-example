package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("hello")

	//secret leak
	aws_secret := "AKIAIMNOJVGFDXXXE4OA"
	logrus.Info(aws_secret)

	aws_secret2 := "AKIAIMNOJVGFDXXXE4OB"
        logrus.Info(aws_secret2)
}
