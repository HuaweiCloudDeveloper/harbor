//go:build include_obs
// +build include_obs

package main

import (
	"errors"
	"testing"

	"github.com/docker/distribution/registry/storage/driver/factory"
)

func TestOBSDriverRegistered(t *testing.T) {
	_, err := factory.Create("obs", map[string]interface{}{})
	if err == nil {
		return
	}

	var invalid factory.InvalidStorageDriverError
	if errors.As(err, &invalid) {
		t.Fatalf("obs driver not registered: %v", err)
	}
}

