package cctx

import (
	"testing"
)

func TestContextSetAndGet(t *testing.T) {
	ctx := NewContext()

	// Test setting and getting a value
	ctx.Set("key", "value")
	val := ctx.Get("key")

	if val != "value" {
		t.Errorf("Expected 'value', but got '%v'", val)
	}
}

func TestContextDelete(t *testing.T) {
	ctx := NewContext()

	// Test deleting a value
	ctx.Set("key", "value")
	ctx.Delete("key")
	val := ctx.Get("key")

	if val != nil {
		t.Errorf("Expected nil, but got '%v'", val)
	}
}

func TestContextGetNonExistentKey(t *testing.T) {
	ctx := NewContext()

	// Test getting a non-existent key
	val := ctx.Get("nonexistent")

	if val != nil {
		t.Errorf("Expected nil, but got '%v'", val)
	}
}

func TestContextDeleteNonExistentKey(t *testing.T) {
	ctx := NewContext()

	// Test deleting a non-existent key
	ctx.Delete("nonexistent")
	// No error expected, should not panic
}
