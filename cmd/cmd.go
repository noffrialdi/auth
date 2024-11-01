package cmd

import (
	"log"
	"os"

	"github.com/noffrialdi/auth/cmd/server"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Auth Service",
		Short: "Auth - Backend Service",
		Long:  "Auth - API Gateway Auth Service",
	}
)

func Execute() {
	rootCmd.AddCommand(server.ServeHTTPCmd())
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Error: \n", err.Error())
		os.Exit(-1)
	}
}
