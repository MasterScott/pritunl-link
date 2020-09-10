package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/pritunl/pritunl-link/clean"
	"github.com/pritunl/pritunl-link/constants"
	"github.com/pritunl/pritunl-link/sync"
	"github.com/pritunl/pritunl-link/watch"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() (err error) {
	logrus.WithFields(logrus.Fields{
		"version": constants.Version,
	}).Info("cmd.start: Starting link")

	sync.Init()
	watch.Init()

	sig := make(chan os.Signal, 2)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	constants.Interrupt = true

	clean.CleanUp()

	time.Sleep(1010 * time.Millisecond)

	clean.CleanUp()

	return
}
