package cmd

import (
	"github.com/jantytgat/go-kit/pkg/application"
)

var SubCommands = []application.Commander{
	acmeCmd,
	daemonCmd,
}
