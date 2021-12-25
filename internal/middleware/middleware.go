package router

import (
	"context"
	"encoding/json"
	"github.com/rcrowley/go-metrics"
	"net/http"
	"strings"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog/hlog"
	"github.com/zigapk/prpo-auth/internal/logger"
	"github.com/zigapk/prpo-auth/internal/models/user"
)

const contextUserKey = "auth_user"

// createLoggerMiddleware creates logger middleware.
func createLoggerMiddleware() func(handler http.Handler) http.Handler {
	return hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		u := UserFromRequest(r)

		var userID string
		if u != nil {
			userID = u.UID
		}

		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Bool("authenticated", u != nil).
			Str("user_id", userID).
			Send()
	})
}

// metricsMiddleware counts number of requests.
func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		meter := metrics.GetOrRegisterMeter("requests", metrics.DefaultRegistry)
		meter.Mark(1)

		next.ServeHTTP(w, r)
	})
}

func writeCorsHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Authorization")
}

func CorsHandler(w http.ResponseWriter, _ *http.Request) {
	writeCorsHeader(w)

	_, _ = w.Write([]byte{})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeCorsHeader(w)
		next.ServeHTTP(w, r)
	})
}

func responseTypeHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func tokenFromRequest(r *http.Request) string {
	// Get header value.
	header := r.Header.Get("Authorization")

	// Split header value to components.
	authComponents := strings.Split(header, " ")
	if len(authComponents) == 2 {
		// Extract auth type and token
		authType := authComponents[0]
		token := authComponents[1]

		// Check if auth type is bearer
		if authType == "Bearer" {
			return token
		}
	}

	// Get token from GET parameter.
	return r.URL.Query().Get("key")
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get access token string from request.
		accessToken := tokenFromRequest(r)

		// Check if token is valid and get user from token.
		u, err := user.FromToken(accessToken)

		// Add user to context if valid.
		if err == nil {
			ctx := context.WithValue(r.Context(), contextUserKey, u)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func AddMiddleware(handler http.Handler, authorizedOnly, customContentType bool) http.Handler {
	c := alice.New(hlog.NewHandler(logger.Log), Middleware, createLoggerMiddleware(), metricsMiddleware)

	if !customContentType {
		c = c.Append(responseTypeHeaderMiddleware)
	}

	c = c.Append(corsMiddleware)

	if authorizedOnly {
		c = c.Append(RequiredMiddleware)
	}

	return c.Then(handler)
}

func UserFromRequest(r *http.Request) *user.User {
	u, ok := r.Context().Value(contextUserKey).(*user.User)
	if ok {
		return u
	} else {
		return nil
	}
}

func RequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := UserFromRequest(r)
		if u == nil {
			// Write error to response.
			res, _ := json.Marshal(map[string]interface{}{
				"error_code": "invalid_token",
				"metadata":   nil,
			})

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write(res)
		} else {
			// Continue to next handler.
			next.ServeHTTP(w, r)
		}
	})
}
