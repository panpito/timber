package main

import log "github.com/sirupsen/logrus"

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
		ForceColors:      true,
	})
}
