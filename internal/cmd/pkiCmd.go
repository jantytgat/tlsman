package cmd

import (
	"github.com/jantytgat/go-kit/pkg/application"
	"github.com/spf13/cobra"
)

const (
	pkiCmdUse   = "pki"
	pkiCmdShort = "Manage PKI"
	pkiCmdLong  = "Manage PKI"
)

var pkiCmd = application.Command{
	Command: &cobra.Command{
		Use:   pkiCmdUse,
		Short: pkiCmdShort,
		Long:  pkiCmdLong,
	},
	SubCommands: []application.Commander{},
	Configure:   nil,
}
