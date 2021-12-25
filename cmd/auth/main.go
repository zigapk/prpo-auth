package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/zigapk/prpo-auth/internal/cmd"
	"github.com/zigapk/prpo-auth/internal/config"
	"github.com/zigapk/prpo-auth/internal/database"
	"github.com/zigapk/prpo-auth/internal/logger"
	token "github.com/zigapk/prpo-auth/internal/util"
)

// @title        PRPO Auth API
// @version      1.0
// @description  PRPO Auth service api.

// @contact.name   Žiga Patačko Koderman
// @contact.url    https://zerodays.dev
// @contact.email  ziga@zerodays.dev
func main() {
	// Load config
	config.Load()

	// Init logger
	logger.Init()

	// Init database.
	database.Init()
	defer database.Close()

	// Set public key for tokens.
	token.SetKey(config.Login.SigningPublicKey)

	// Create cli app.
	app := cli.NewApp()
	app.Name = "PRPO auth microservice."
	app.Version = "0.0.1"
	app.Authors = []*cli.Author{{Name: "Žiga Patačko Koderman", Email: "ziga.patacko@gmail.com"}}

	app.Commands = []*cli.Command{
		cmd.Serve,
		cmd.GenKeys,
		cmd.CreateUser,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
