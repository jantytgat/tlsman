package main

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jantytgat/go-kit/pkg/application"
	"github.com/jantytgat/go-kit/pkg/semver"
	"github.com/jantytgat/go-kit/pkg/slogd"
	"github.com/jantytgat/go-registry/pkg/registry"
	"github.com/jantytgat/go-sql-queryrepo/pkg/queryrepo"
	"github.com/spf13/cobra"

	"github.com/jantytgat/tlsman/internal/cmd"
)

//go:generate sh -c "cp -R ../../external/registry/assets ."

var (
	version string

	appName   = "tlsman"
	appTitle  = "TLSMAN"
	appBanner = ""
)

var (
	semVersion semver.Version
	ctx        context.Context
)

var (
	db   *sql.DB
	rsql *queryrepo.Repository
)

//go:embed assets/migrations/*
var migrationFS embed.FS

//go:embed assets/queries/*
var queryFS embed.FS

func main() {
	var err error
	slogd.RegisterColoredTextHandler(os.Stdout, true)

	if semVersion, err = semver.Parse(version); err != nil {
		slogd.Logger().LogAttrs(ctx, slogd.LevelError, "error running application", slog.Any("error", err))
	}

	ctx = slogd.WithContext(context.Background())

	application.New(appName, appTitle, appBanner, semVersion)
	application.RegisterCommands(cmd.SubCommands, nil)

	// Set up database
	application.RegisterPersistentPreRunE(initializeDb)
	application.RegisterPersistentPostRunE(checkDb)

	if err = application.Run(ctx); err != nil {
		slogd.Logger().LogAttrs(ctx, slogd.LevelError, "error running application", slog.Any("error", err))
		os.Exit(1)
	}
	os.Exit(0)
}

func initializeDb(cmd *cobra.Command, args []string) error {
	var err error

	if err = initializeInMemory(ctx); err != nil {
		slogd.FromContext(ctx).LogAttrs(ctx, slogd.LevelError, "error initializing database in memory", slog.Any("error", err))
	}

	if err = migrateDb(ctx); err != nil {
		slogd.FromContext(ctx).LogAttrs(ctx, slogd.LevelError, "error migrating database", slog.Any("error", err))
	}

	// Create query repository from embedded files
	if rsql, err = queryrepo.NewFromFs(queryFS, "assets/queries"); err != nil {
		slogd.FromContext(ctx).LogAttrs(ctx, slogd.LevelError, "error initializing query repository", slog.Any("error", err))
	}
	return nil
}

func checkDb(cmd *cobra.Command, args []string) error {
	var err error
	var organizations []registry.Organization
	if organizations, err = registry.ListOrganizations(rsql, db); err != nil {
		return err
	}

	for _, o := range organizations {
		fmt.Println(o)
	}

	return nil
}

func initializeInMemory(ctx context.Context) error {
	var err error
	db, err = sql.Open("sqlite", ":memory:")
	if err != nil {
		slogd.FromContext(ctx).LogAttrs(ctx, slogd.LevelError, "error initializing database connection", slog.Any("error", err))
		return err
	}
	return nil
}

func migrateDb(ctx context.Context) error {
	var err error
	var src source.Driver
	if src, err = iofs.New(migrationFS, "assets/migrations"); err != nil {
		slogd.FromContext(ctx).LogAttrs(ctx, slogd.LevelError, "failed to load migrations filesystem", slog.Any("error", err))
		return err
	}

	var driver database.Driver
	if driver, err = sqlite.WithInstance(db, &sqlite.Config{}); err != nil {
		slogd.FromContext(ctx).LogAttrs(ctx, slogd.LevelError, "migrations target connection failed", slog.Any("error", err))
		return err
	}

	var m *migrate.Migrate
	if m, err = migrate.NewWithInstance("fs", src, "sqlite", driver); err != nil {
		slogd.FromContext(ctx).LogAttrs(ctx, slogd.LevelError, "migration initialization failed", slog.Any("error", err))
		return err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		slogd.FromContext(ctx).LogAttrs(ctx, slogd.LevelError, "migration failed", slog.Any("error", err))
		return err
	}
	return nil
}
