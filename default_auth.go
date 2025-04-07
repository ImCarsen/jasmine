package jasmine

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

// Default middleware struct for the auth middleware. This is just an example and should be replaced with your own implementation.
type DefaultAuthMiddleware struct {
	Next http.Handler
}

func (DefaultAuthMiddleware) GetName() string {
	return "DefaultAuthMiddleware"
}
func (DefaultAuthMiddleware) GetDescription() string {
	return "Default Auth Middleware"
}
func (m DefaultAuthMiddleware) GetNext() http.Handler {
	return m.Next
}
func (m DefaultAuthMiddleware) GetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("Protected route accessed")
		// Example for if the user is not authorized
		if !true {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Call the next handler
		m.Next.ServeHTTP(w, r)
	})
}

// Default/Example AuthFunc. This doesn't actually do anything, you will want to create your own func for handling protected routes.
func DefaultAuthFunc(next http.Handler) Middleware {
	return &DefaultAuthMiddleware{
		Next: next,
	}
}
