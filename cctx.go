package cctx

// Context is the main interface that combines ContextReader, ContentWriter, and ContextRemover.
type Context interface {
	ContextWriter
	ContextReader
	ContextRemover
}

// ContextReader is the interface for reading values from the context.
type ContextReader interface {
	// Get retrieves a value from the context by the given key.
	Get(key string) interface{}
}

// ContentWriter is the interface for writing values to the context.
type ContextWriter interface {
	// Set sets a value in the context with the given key.
	Set(key string, value interface{})
}

// ContextRemover is the interface for removing values from the context.
type ContextRemover interface {
	// Delete removes a value from the context by the given key.
	Delete(key string)
}

// context is the default implementation of the Context interface.
type context struct {
	data map[string]interface{}
}

// NewContext creates a new instance of the context.
func NewContext() Context {
	return &context{
		data: make(map[string]interface{}),
	}
}

func (c *context) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *context) Get(key string) interface{} {
	return c.data[key]
}

func (c *context) Delete(key string) {
	delete(c.data, key)
}
