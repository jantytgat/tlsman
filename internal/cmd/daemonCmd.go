package cmd

import (
	"context"
	"net/http"
	"sync"

	"github.com/jantytgat/go-kit/pkg/application"
	"github.com/jantytgat/go-kit/pkg/slogd"
	"github.com/spf13/cobra"
)

const (
	daemonCmdUse   = "daemon"
	daemonCmdShort = "Run daemon"
	daemonCmdLong  = "Run daemon"
)

var daemonCmd = application.Command{
	Command: &cobra.Command{
		Use:   daemonCmdUse,
		Short: daemonCmdShort,
		Long:  daemonCmdLong,
		RunE:  daemonRunE,
	},
	SubCommands: nil,
	Configure:   nil,
}

func daemonRunE(cmd *cobra.Command, args []string) error {
	slogd.FromContext(cmd.Context()).LogAttrs(cmd.Context(), slogd.LevelInfo, cmd.CommandPath())
	mux := http.NewServeMux() // Create sample handler to returns 404
	mux.Handle("/", http.RedirectHandler("https://jantytgat.com", 302))

	srv := application.NewSocketHttpServer("tlsman.socket", mux)
	srvCtx, srvCancel := context.WithCancel(cmd.Context())
	defer srvCancel()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		srv.Run(ctx)
	}(srvCtx, &wg)

	wg.Wait()
	return nil
}
