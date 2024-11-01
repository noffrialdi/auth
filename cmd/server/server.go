package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/noffrialdi/auth/config"
	"github.com/noffrialdi/auth/internal/handler"
	"github.com/noffrialdi/auth/internal/infrastructures/database"
	"github.com/spf13/cobra"
)

var (
	serveHTTPCmd = &cobra.Command{
		Use:   "serve-http",
		Short: "CRUD Service Service HTTP",
		Long:  "Serve CRUD Service Service through HTTP",
		RunE:  run,
	}
)

func ServeHTTPCmd() *cobra.Command {
	serveHTTPCmd.Flags().StringP("config", "c", "", "Config Path, both relative or absolute. i.e: /usr/local/bin/config/files")
	return serveHTTPCmd
}

func run(cmd *cobra.Command, args []string) error {
	// ctx := context.Background()
	configLocation, err := cmd.Flags().GetString("config")
	if err != nil {
		log.Fatalf("failed to connect redis err:%v", err.Error())
	}

	cfg := &config.MainConfig{}
	config.ReadModuleConfig(cfg, "main", configLocation)

	dataStore := database.New(&database.DatabaseConfig{
		MasterDSN:       cfg.Database.MasterDSN,
		MaxIdleConn:     cfg.Database.MaxIdleConn,
		MaxConn:         cfg.Database.MaxConn,
		ConnMaxLifetime: cfg.Database.ConnMaxLifetime,
	})

	appContainer := newContainer(&Opts{
		Cfg:          cfg,
		MasterDataDB: dataStore,
	})

	server := handler.NewHTTP(&handler.Opts{
		Cfg:  appContainer.Cfg,
		User: appContainer.Auth,
	})

	log.Println("Server is starting ... ")
	go server.Run()

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Println("Exiting gracefully...")
	case err := <-server.ListenError():
		log.Fatalln("Error starting web server, exiting gracefully:", err)
	}

	return nil
}
