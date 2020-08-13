package main

import (
	"github.com/google/gops/agent"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := agent.Listen(agent.Options{}); err != nil {
		logrus.Fatal(err)
	}
	StartServerWithDefaults()
}
