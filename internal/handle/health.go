package handle

import (
	"encoding/json"
	"github.com/rcrowley/go-metrics"
	"github.com/zigapk/prpo-auth/internal/database"
	"net/http"
)

type livenessResponse struct {
	Status string `json:"status"`
}

type readinessResponse struct {
	Status   string `json:"status"`
	DBStatus string `json:"db_status"`
}

// LivenessHandle  @Summary      Liveness probe
// @Description    Get liveness status from service.
// @Produce        json
// @Success        200  {object}  livenessResponse
// @Router         /health/liveness [get]
func LivenessHandle(w http.ResponseWriter, _ *http.Request) {
	resp := livenessResponse{}
	resp.Status = "UP"
	w.WriteHeader(http.StatusOK)

	res, _ := json.Marshal(resp)
	_, _ = w.Write(res)
}

// ReadinessHandle  @Summary      Readiness probe
// @Description     Returns weather microservice is ready to accept connections.
// @Produce         json
// @Success         200  {object}  readinessResponse
// @Router          /health/readiness [get]
func ReadinessHandle(w http.ResponseWriter, _ *http.Request) {
	resp := readinessResponse{}
	ok := true

	// Check if database is ready.
	dbErr := database.DB.Ping()
	if dbErr == nil {
		resp.DBStatus = "UP"
	} else {
		ok = false
		resp.DBStatus = "DOWN: " + dbErr.Error()
	}

	// Set global status
	if ok {
		resp.Status = "UP"
		w.WriteHeader(http.StatusOK)
	} else {
		resp.Status = "DOWN"
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	res, _ := json.Marshal(resp)
	_, _ = w.Write(res)
}

// MetricsHandle  @Summary      Metrics
// @Description   Returns microservice metrics data.
// @Produce       json
// @Router        /metrics [get]
func MetricsHandle(w http.ResponseWriter, _ *http.Request) {
	metrics.WriteJSONOnce(metrics.DefaultRegistry, w)
}
