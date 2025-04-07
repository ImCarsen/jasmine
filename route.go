package jasmine

import "net/http"

type RouteHandler interface {
	Name() string
	Description() string
	Category() string
	Handler() http.HandlerFunc
}
