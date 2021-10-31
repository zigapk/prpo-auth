package errors

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, err ResponseError) {
	res, _ := json.Marshal(map[string]interface{}{
		"error_code": err.Code,
		"metadata":   err.Metadata,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	_, _ = w.Write(res)
}
