package main

import (
	"bettersocial/config"
	"bettersocial/container"
	"bettersocial/repository/gormrepo"
	"bettersocial/storage"

	"context"
	"os"
	_ "time/tzdata"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var rootCmd = &cobra.Command{
	Use: "bettersocial",
}

func init() {
	loadConfig()
}

func main() {
	Execute()
}

func Execute() {
	rootCmd := registerCommands(&defaultAppProvider{})
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err.Error())
		os.Exit(1)
	}
}

func loadConfig() {
	err := config.Load()
	if err != nil {
		logrus.Errorf("Config error: %s", err.Error())
		os.Exit(1)
	}
}

func registerCommands(appProvider AppProvider) *cobra.Command {
	rootCmd.AddCommand(Server(appProvider))
	rootCmd.AddCommand(Migrate(appProvider))

	return rootCmd
}

type AppProvider interface {
	BuildContainer(ctx context.Context, options buildOptions) (*container.Container, func(), error)
}

type buildOptions struct {
	Postgres bool
}

type defaultAppProvider struct {
}

func (defaultAppProvider) BuildContainer(ctx context.Context, options buildOptions) (*container.Container, func(), error) {
	var db *gorm.DB
	cfg := config.Instance()

	appContainer := container.NewContainer()
	appContainer.SetConfig(cfg)

	if options.Postgres {
		db = storage.GetPostgresDb()
		appContainer.SetDb(db)

		todoRepo := gormrepo.NewUserRepository(db)
		imageRepo := gormrepo.NewImageRepository(db)
		appContainer.SetUserRepo(todoRepo)
		appContainer.SetImageRepo(imageRepo)
	}

	deferFn := func() {
		if db != nil {
			storage.CloseDB(db)
		}
	}

	return appContainer, deferFn, nil
}
