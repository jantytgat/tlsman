package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jantytgat/go-kit/pkg/application"
	"github.com/jantytgat/go-kit/pkg/semver"
	"github.com/jantytgat/go-kit/pkg/slogd"
	"github.com/spf13/cobra"
)

var (
	BuildVersion string

	appName   = "tlsman"
	appTitle  = "TLSMAN"
	appBanner = ""
)

var (
	version semver.Version
	ctx     context.Context

	tlsmanCmd = &cobra.Command{
		Use:  appName,
		RunE: tlsmanFuncE,
	}
)

func tlsmanFuncE(cmd *cobra.Command, args []string) error {
	return nil
}

func main() {
	var err error

	slogd.Init(slogd.LevelInfo, false)
	slogd.RegisterColoredTextHandler(os.Stderr, true)

	if version, err = semver.Parse(BuildVersion); err != nil {
		slogd.Logger().LogAttrs(ctx, slogd.LevelError, "error running application", slog.Any("error", err))
	}

	ctx = slogd.WithContext(context.Background())

	application.New(appName, appTitle, appBanner, version)
	application.RegisterCommand(tlsmanCmd)

	if err = application.Run(ctx); err != nil {
		slogd.Logger().LogAttrs(ctx, slogd.LevelError, "error running application", slog.Any("error", err))
		os.Exit(1)
	}
	os.Exit(0)
}
