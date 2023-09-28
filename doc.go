// Package cctx provides a context management package with read, write, and delete operations.
//
// This package defines the Context interface, which combines three sub-interfaces:
// - ContextReader: For reading values from the context.
// - ContentWriter: For writing values to the context.
// - ContextRemover: For removing values from the context.
//
// The default implementation of the Context interface is provided by the context struct, which
// can be created using the NewContext function.
//
// Example Usage:
//
//	// Create a new context
//	ctx := cctx.NewContext()
//
//	// Set a value in the context
//	ctx.Set("key", "value")
//
//	// Retrieve a value from the context
//	val := ctx.Get("key")
//
//	// Delete a value from the context
//	ctx.Delete("key")
//
// This package allows you to manage contextual data within your application.
package cctx
