package main

import (
	"github.com/fr-str/log"
)

func main() {
	log.Debug("dupa", log.String("f1", "dupa2"))
	log.Info("dupa")
	log.Warn("dupa")
	log.Error("dupa")
}
