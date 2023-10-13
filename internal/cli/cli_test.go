package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.NotPanics(t, func() { Run("0.0.0", []string{"go-ebsnvme", "--version"}) })
}

func TestNewApp(t *testing.T) {
	app := NewApp("0.0.0")
	assert.Equal(t, "go-ebsnvme", app.Name)
	assert.Equal(t, "0.0.0", app.Version)
}
