package api_test

import (
	"media/internal/api"
	"testing"
)

func TestRouterInitializes(t *testing.T) {
	if err := api.CreateRouter(); err == nil {
		t.Errorf("Expected router to not be nil")
	}
}
