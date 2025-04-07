package jasmine

import "net/http"

// Allows the creation of middleware handlers dynamically.
func NewMiddleware(name, description, category string, next http.Handler, handlerFunc func() http.HandlerFunc) Middleware {
	return &ConstructableMiddleware{
		MName:          name,
		MDescription:   description,
		MCategory:      category,
		NextHandler:    next,
		GetHandlerFunc: handlerFunc,
	}
}

// ConstructableMiddleware is a struct that implements the Middleware interface. It is used to create middleware handlers dynamically.
type ConstructableMiddleware struct {
	MName          string                  // Name of the middleware
	MDescription   string                  // Description of the middleware
	MCategory      string                  // Category of the middleware
	NextHandler    http.Handler            // The next handler in the chain
	GetHandlerFunc func() http.HandlerFunc // Function to get the handler function
}

func (c *ConstructableMiddleware) Name() string {
	return c.MName
}

func (c *ConstructableMiddleware) Description() string {
	return c.MDescription
}

func (c *ConstructableMiddleware) Category() string {
	return c.MCategory
}

func (c *ConstructableMiddleware) Next() http.Handler {
	return c.NextHandler
}

func (c *ConstructableMiddleware) Handler() http.Handler {
	return c.GetHandlerFunc()
}
