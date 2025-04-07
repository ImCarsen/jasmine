package jasmine

import (
	"net/http"

	"github.com/rs/zerolog"
)

type AuthFunc func(http.Handler) Middleware

type Routes struct {
	Routes          map[string]RouteHandler // path -> handler | Unprotected
	ProtectedRoutes map[string]RouteHandler // path -> handler | Auto registers behind the AuthFunc
	AuthFunc        AuthFunc                // func(http.HandlerFunc) http.Handler | ProtectedRoutes will not work without this
}

// Registers all routes supplied
func (s *Routes) RegisterRoutes(mux *http.ServeMux, logger *zerolog.Logger) {
	// - Unprotected routes
	// Make sure there is at least 1 route defined
	if len(s.Routes) == 0 {
		logger.Error().Msg("No routes provided. Please add at least one route")
		return
	}
	// Register the handler with the muxer
	for path, handler := range s.Routes {
		//logger.Debug().Str("route", path).Msg("Registering unprotected route")
		mux.HandleFunc(path, handler.Handler())
	}
	// - Protected routes
	// Make sure the auth func is set for protected routes
	if s.AuthFunc == nil {
		logger.Info().Msg("Auth function for protected routes not set, skipping.")
		return
	}
	// Make sure there are protected routes, if not, return
	if len(s.ProtectedRoutes) == 0 {
		logger.Info().Msg("No protected routes defined, skipping.")
		return
	}
	// Register the handler with the muxer wrapped by the AuthFunc
	for path, handler := range s.ProtectedRoutes {
		//logger.Debug().Str("route", path).Msg("Registering protected route")
		mux.Handle(path, s.AuthFunc(handler.Handler()).Handler())
	}
}
