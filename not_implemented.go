package jasmine

import "net/http"

// NotImplemented RouteHandler

// The default handler for routes that are not implemented
var NotImplemented = &NotImplementedRoute{}

type NotImplementedRoute struct {
}

func (n *NotImplementedRoute) GetName() string {
	return "Not Implemented"
}

func (n *NotImplementedRoute) GetDescription() string {
	return "This route is not implemented yet"
}

func (n *NotImplementedRoute) GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Implemented", http.StatusNotImplemented)
	}
}
