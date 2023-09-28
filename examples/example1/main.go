package main

import (
	"fmt"

	"github.com/ssapp/cctx"
)

func main() {
	// Create a new context
	ctx := cctx.NewContext()

	// Set a value in the context
	ctx.Set("key", "value")

	// Retrieve a value from the context
	val := ctx.Get("key")

	// Delete a value from the context
	ctx.Delete("key")

	fmt.Println(val)
}
