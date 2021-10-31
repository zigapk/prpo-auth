package util

import (
	"encoding/json"
	"net/http"

	"github.com/zigapk/prpo-auth/internal/handle/errors"
	"github.com/zigapk/prpo-auth/internal/logger"
)

func ParseJSON(w http.ResponseWriter, r *http.Request, dest interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dest)
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		errors.Response(w, errors.InvalidJSON)

		return false
	}

	return true
}
