package handle

import (
	"net/http"

	"github.com/zigapk/prpo-auth/internal/config"
	"github.com/zigapk/prpo-auth/internal/logger"
	"github.com/zigapk/prpo-auth/internal/util"
)

func SigningKeyHandle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	err := util.EncodePubKey(config.Login.SigningPublicKey, w)
	if err != nil {
		logger.Log.Warn().Err(err).Send()
	}
}
