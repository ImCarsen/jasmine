package jasmine

import (
	"net/http"
)

// Middleware interface defines the methods that a middleware must implement.
type Middleware interface {
	Name() string
	Description() string
	Category() string
	Next() http.Handler
	Handler() http.Handler
}
