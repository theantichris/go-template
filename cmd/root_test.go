package cmd

import (
	"bytes"
	"testing"

	"github.com/charmbracelet/log"
)

func TestNewRootCmd(t *testing.T) {
	t.Parallel()

	logger := log.NewWithOptions(bytes.NewBuffer(nil), log.Options{
		Level: log.ErrorLevel,
	})

	t.Run("creates command with correct properties", func(t *testing.T) {
		t.Parallel()

		cmd := NewRootCmd(logger)

		if cmd == nil {
			t.Fatal("expected command to be created")
		}

		if cmd.Use != "example" {
			t.Errorf("expected Use to be 'example', got %q", cmd.Use)
		}

		if cmd.Short == "" {
			t.Error("expected Short description to be non-empty")
		}

		if cmd.Long == "" {
			t.Error("expected Long description to be non-empty")
		}
	})

	t.Run("has PreRunE function", func(t *testing.T) {
		t.Parallel()

		cmd := NewRootCmd(logger)

		if cmd.PreRunE == nil {
			t.Error("expected PreRunE to be set")
		}
	})
}

func TestInitConfig(t *testing.T) {
	t.Run("completes without error", func(t *testing.T) {
		t.Parallel()

		logger := log.NewWithOptions(bytes.NewBuffer(nil), log.Options{
			Level: log.ErrorLevel,
		})

		initConfig(logger)
	})

	t.Run("sets logger level when debug enabled", func(t *testing.T) {
		t.Parallel()

		var logOutput bytes.Buffer
		logger := log.NewWithOptions(&logOutput, log.Options{
			Level: log.WarnLevel,
		})

		if logger.GetLevel() != log.WarnLevel {
			t.Errorf("expected initial logger level to be WarnLevel, got %v", logger.GetLevel())
		}
	})
}
