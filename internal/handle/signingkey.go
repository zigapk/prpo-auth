package handle

import (
	"net/http"

	"github.com/zigapk/prpo-auth/internal/config"
	"github.com/zigapk/prpo-auth/internal/logger"
	"github.com/zigapk/prpo-auth/internal/util"
)

// SigningKeyHandle  @Summary      Get the JWT signing key.
// @Description      Get the JWT signing key.
// @Produce          application/text
// @Success          200
// @Failure          500  {object}  errors.ResponseError
// @Router           /signing_key/ [get]
func SigningKeyHandle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	err := util.EncodePubKey(config.Login.SigningPublicKey, w)
	if err != nil {
		logger.Log.Warn().Err(err).Send()
	}
}
