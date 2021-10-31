package cmd

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
	"github.com/zigapk/prpo-auth/internal/config"
	"github.com/zigapk/prpo-auth/internal/logger"
	"github.com/zigapk/prpo-auth/internal/router"
)

var Serve = &cli.Command{
	Name:   "serve",
	Usage:  "Start server.",
	Action: serve,
}

func serve(_ *cli.Context) error {
	// Check that signing keys are valid.
	if config.Login.SigningPublicKey == nil || config.Login.SigningPrivateKey == nil {
		logger.Log.Fatal().Msg("Can not run server without valid private and public signing keys.")
	}

	// Load router.
	r := router.NewRouter()

	// Start listening for connections.
	listenAddress := fmt.Sprintf("%s:%d", config.Server.ListenAddress(), config.Server.Port())
	logger.Log.Fatal().Err(http.ListenAndServe(listenAddress, r)).Send()

	return nil
}
