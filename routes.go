package jasmine

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AuthFunc func(http.HandlerFunc) http.Handler

type Routes struct {
	Routes          map[string]http.HandlerFunc // path -> handler | Unprotected
	ProtectedRoutes map[string]http.HandlerFunc // path -> handler | Auto registers behind the AuthFunc
	AuthFunc        AuthFunc                    // func(http.HandlerFunc) http.Handler | ProtectedRoutes will not work without this
}

// Registers all routes supplied
func (s *Routes) RegisterRoutes(mux *http.ServeMux, logger *zerolog.Logger) {
	// - Unprotected routes
	// Make sure there is at least 1 route defined
	if s.Routes == nil || len(s.Routes) == 0 {
		logger.Error().Msg("No routes provided. Please add at least one route")
		return
	}
	// Register the handler with the muxer
	for path, handler := range s.Routes {
		mux.HandleFunc(path, handler)
	}
	// - Protected routes
	// Make sure the auth func is set for protected routes
	if s.AuthFunc == nil {
		logger.Info().Msg("Auth function for protected routes not set, skipping.")
		return
	}
	// Make sure there are protected routes, if not, return
	if s.ProtectedRoutes == nil || len(s.ProtectedRoutes) == 0 {
		logger.Info().Msg("No protected routes defined, skipping.")
		return
	}
	// Register the handler with the muxer wrapped by the AuthFunc
	for path, handler := range s.ProtectedRoutes {
		mux.Handle(path, s.AuthFunc(handler))
	}
}

// Default/Example AuthFunc. This doesn't actually do anything, you will want to create your own func for handling protected routes.
func DefaultAuthFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("Protected route accessed")
		// Example for if the user is not authorized
		if !true {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
