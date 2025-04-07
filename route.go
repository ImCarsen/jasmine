package jasmine

import "net/http"

type RouteHandler interface {
	GetName() string
	GetDescription() string
	GetHandler() http.HandlerFunc
}
