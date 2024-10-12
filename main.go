package main

import (
	"github.com/hwdef/topology/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	cmd.Execute()
}
