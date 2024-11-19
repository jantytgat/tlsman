package cmd

import (
	"github.com/jantytgat/go-kit/pkg/application"
	"github.com/jantytgat/go-kit/pkg/slogd"
	"github.com/spf13/cobra"
)

const (
	acmeCmdUse   = "acme"
	acmeCmdShort = "Manage ACME integration"
	acmeCmdLong  = "Manage ACME integration"

	acmeAccountCmdUse   = "account"
	acmeAccountCmdShort = "Manage ACME account"
	acmeAccountCmdLong  = "Manage ACME account"

	acmeAccountAddCmdUse   = "add"
	acmeAccountAddCmdShort = "Add ACME account"
	acmeAccountAddCmdLong  = "Add ACME account"

	acmeAccountDeleteCmdUse   = "delete"
	acmeAccountDeleteCmdShort = "Delete ACME account"
	acmeAccountDeleteCmdLong  = "Delete ACME account"

	acmeAccountListCmdUse   = "list"
	acmeAccountListCmdShort = "List ACME accounts"
	acmeAccountListCmdLong  = "List ACME accounts"

	acmeAccountUpdateCmdUse   = "update"
	acmeAccountUpdateCmdShort = "Update ACME account"
	acmeAccountUpdateCmdLong  = "Update ACME account"

	acmeServiceCmdUse   = "service"
	acmeServiceCmdShort = "Manage ACME service"
	acmeServiceCmdLong  = "Manage ACME service"

	acmeServiceAddCmdUse   = "add"
	acmeServiceAddCmdShort = "Add ACME service"
	acmeServiceAddCmdLong  = "Add ACME service"

	acmeServiceDeleteCmdUse   = "delete"
	acmeServiceDeleteCmdShort = "Delete ACME service"
	acmeServiceDeleteCmdLong  = "Delete ACME service"

	acmeServiceListCmdUse   = "list"
	acmeServiceListCmdShort = "List ACME service"
	acmeServiceListCmdLong  = "List ACME service"

	acmeServiceUpdateCmdUse   = "update"
	acmeServiceUpdateCmdShort = "Update ACME service"
	acmeServiceUpdateCmdLong  = "Update ACME service"

	acmeProviderCmdUse   = "provider"
	acmeProviderCmdShort = "Manage ACME provider"
	acmeProviderCmdLong  = "Manage ACME provider"

	acmeProviderAddCmdUse   = "add"
	acmeProviderAddCmdShort = "Add ACME provider"
	acmeProviderAddCmdLong  = "Add ACME provider"

	acmeProviderDeleteCmdUse   = "delete"
	acmeProviderDeleteCmdShort = "Delete ACME provider"
	acmeProviderDeleteCmdLong  = "Delete ACME provider"

	acmeProviderListCmdUse   = "list"
	acmeProviderListCmdShort = "List ACME provider"
	acmeProviderListCmdLong  = "List ACME provider"

	acmeProviderUpdateCmdUse   = "update"
	acmeProviderUpdateCmdShort = "Update ACME provider"
	acmeProviderUpdateCmdLong  = "Update ACME provider"
)

var acmeCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeCmdUse,
		Short: acmeCmdShort,
		Long:  acmeCmdLong,
	},
	SubCommands: []application.Commander{
		acmeAccountCmd,
		acmeServiceCmd,
		acmeProviderCmd,
	},
	Configure: nil,
}

var acmeAccountCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeAccountCmdUse,
		Short: acmeAccountCmdShort,
		Long:  acmeAccountCmdLong,
	},
	SubCommands: []application.Commander{
		acmeAccountAddCmd,
		acmeAccountDeleteCmd,
		acmeAccountListCmd,
		acmeAccountUpdateCmd,
	},
	Configure: nil,
}

var acmeAccountAddCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeAccountAddCmdUse,
		Short: acmeAccountAddCmdShort,
		Long:  acmeAccountAddCmdLong,
		RunE:  acmeAccountAddRunE,
	},
	SubCommands: nil,
	Configure:   nil,
}

func acmeAccountAddRunE(cmd *cobra.Command, args []string) error {
	slogd.FromContext(cmd.Context()).LogAttrs(cmd.Context(), slogd.LevelInfo, cmd.CommandPath())
	return nil
}

var acmeAccountDeleteCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeAccountDeleteCmdUse,
		Short: acmeAccountDeleteCmdShort,
		Long:  acmeAccountDeleteCmdLong,
		RunE:  acmeAccountDeleteRunE,
	},
	SubCommands: nil,
	Configure:   nil,
}

func acmeAccountDeleteRunE(cmd *cobra.Command, args []string) error {
	slogd.FromContext(cmd.Context()).LogAttrs(cmd.Context(), slogd.LevelInfo, cmd.CommandPath())
	return nil
}

var acmeAccountListCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeAccountListCmdUse,
		Short: acmeAccountListCmdShort,
		Long:  acmeAccountListCmdLong,
		RunE:  acmeAccountListRunE,
	},
	SubCommands: nil,
	Configure:   nil,
}

func acmeAccountListRunE(cmd *cobra.Command, args []string) error {
	slogd.FromContext(cmd.Context()).LogAttrs(cmd.Context(), slogd.LevelInfo, cmd.CommandPath())
	return nil
}

var acmeAccountUpdateCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeAccountUpdateCmdUse,
		Short: acmeAccountUpdateCmdShort,
		Long:  acmeAccountUpdateCmdLong,
		RunE:  acmeAccountUpdateRunE,
	},
	SubCommands: nil,
	Configure:   nil,
}

func acmeAccountUpdateRunE(cmd *cobra.Command, args []string) error {
	slogd.FromContext(cmd.Context()).LogAttrs(cmd.Context(), slogd.LevelInfo, cmd.CommandPath())
	return nil
}

var acmeServiceCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeServiceCmdUse,
		Short: acmeServiceCmdShort,
		Long:  acmeServiceCmdLong,
	},
	SubCommands: []application.Commander{
		acmeServiceAddCmd,
		acmeServiceDeleteCmd,
		acmeServiceListCmd,
		acmeServiceUpdateCmd,
	},
	Configure: nil,
}

var acmeServiceAddCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeServiceAddCmdUse,
		Short: acmeServiceAddCmdShort,
		Long:  acmeServiceAddCmdLong,
		RunE:  acmeServiceAddRunE,
	},
	SubCommands: nil,
	Configure:   nil,
}

func acmeServiceAddRunE(cmd *cobra.Command, args []string) error {
	slogd.FromContext(cmd.Context()).LogAttrs(cmd.Context(), slogd.LevelInfo, cmd.CommandPath())
	return nil
}

var acmeServiceDeleteCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeServiceDeleteCmdUse,
		Short: acmeServiceDeleteCmdShort,
		Long:  acmeServiceDeleteCmdLong,
		RunE:  acmeServiceDeleteRunE,
	},
	SubCommands: nil,
	Configure:   nil,
}

func acmeServiceDeleteRunE(cmd *cobra.Command, args []string) error {
	slogd.FromContext(cmd.Context()).LogAttrs(cmd.Context(), slogd.LevelInfo, cmd.CommandPath())
	return nil
}

var acmeServiceListCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeServiceListCmdUse,
		Short: acmeServiceListCmdShort,
		Long:  acmeServiceListCmdLong,
		RunE:  acmeServiceListRunE,
	},
	SubCommands: nil,
	Configure:   nil,
}

func acmeServiceListRunE(cmd *cobra.Command, args []string) error {
	slogd.FromContext(cmd.Context()).LogAttrs(cmd.Context(), slogd.LevelInfo, cmd.CommandPath())
	return nil
}

var acmeServiceUpdateCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeServiceUpdateCmdUse,
		Short: acmeServiceUpdateCmdShort,
		Long:  acmeServiceUpdateCmdLong,
		RunE:  acmeServiceUpdateRunE,
	},
	SubCommands: nil,
	Configure:   nil,
}

func acmeServiceUpdateRunE(cmd *cobra.Command, args []string) error {
	slogd.FromContext(cmd.Context()).LogAttrs(cmd.Context(), slogd.LevelInfo, cmd.CommandPath())
	return nil
}

var acmeProviderCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeProviderCmdUse,
		Short: acmeProviderCmdShort,
		Long:  acmeProviderCmdLong,
	},
	SubCommands: []application.Commander{
		acmeProviderAddCmd,
		acmeProviderDeleteCmd,
		acmeProviderListCmd,
		acmeProviderUpdateCmd,
	},
	Configure: nil,
}

var acmeProviderAddCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeProviderAddCmdUse,
		Short: acmeProviderAddCmdShort,
		Long:  acmeProviderAddCmdLong,
		RunE:  acmeProviderAddRunE,
	},
	SubCommands: nil,
	Configure:   nil,
}

func acmeProviderAddRunE(cmd *cobra.Command, args []string) error {
	slogd.FromContext(cmd.Context()).LogAttrs(cmd.Context(), slogd.LevelInfo, cmd.CommandPath())
	return nil
}

var acmeProviderDeleteCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeProviderDeleteCmdUse,
		Short: acmeProviderDeleteCmdShort,
		Long:  acmeProviderDeleteCmdLong,
		RunE:  acmeProviderDeleteRunE,
	},
	SubCommands: nil,
	Configure:   nil,
}

func acmeProviderDeleteRunE(cmd *cobra.Command, args []string) error {
	slogd.FromContext(cmd.Context()).LogAttrs(cmd.Context(), slogd.LevelInfo, cmd.CommandPath())
	return nil
}

var acmeProviderListCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeProviderListCmdUse,
		Short: acmeProviderListCmdShort,
		Long:  acmeProviderListCmdLong,
		RunE:  acmeProviderListRunE,
	},
	SubCommands: nil,
	Configure:   nil,
}

func acmeProviderListRunE(cmd *cobra.Command, args []string) error {
	slogd.FromContext(cmd.Context()).LogAttrs(cmd.Context(), slogd.LevelInfo, cmd.CommandPath())
	return nil
}

var acmeProviderUpdateCmd = application.Command{
	Command: &cobra.Command{
		Use:   acmeProviderUpdateCmdUse,
		Short: acmeProviderUpdateCmdShort,
		Long:  acmeProviderUpdateCmdLong,
		RunE:  acmeProviderUpdateRunE,
	},
	SubCommands: nil,
	Configure:   nil,
}

func acmeProviderUpdateRunE(cmd *cobra.Command, args []string) error {
	slogd.FromContext(cmd.Context()).LogAttrs(cmd.Context(), slogd.LevelInfo, cmd.CommandPath())
	return nil
}
