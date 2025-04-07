package jasmine

import "net/http"

type Middleware interface {
	GetName() string
	GetDescription() string
	GetNext() http.Handler
	GetHandler() http.Handler
}
