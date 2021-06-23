package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsolation(t *testing.T) {
	isolationMarker := "/tmp/cisco-datacore-takehome"

	assert.NoFileExists(t, isolationMarker, "Build not properly isolated")

	_, err := os.Create(isolationMarker)
	assert.FileExistsf(t, isolationMarker, "Unable to create build isolation marker: %v", err)
}
