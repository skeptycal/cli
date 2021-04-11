package cli

import log "github.com/sirupsen/logrus"

type logger interface {
}

func NewLogger() logger {
	return log.New()
}
