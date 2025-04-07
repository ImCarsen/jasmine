package jasmine

import "net/http"

type Middleware interface {
	GetName() string
	GetDescription() string
	GetHandler() http.Handler
}
