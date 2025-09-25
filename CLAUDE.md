# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A Go CLI/TUI template project using the Cobra command framework with Charmbracelet tools for terminal UI.

## Common Commands

### Build
```bash
go build
```

### Run
```bash
go run main.go
# Or after building:
./go-template
```

### Test
```bash
go test ./...
go test -v ./...  # verbose output
```

### Module Management
```bash
go mod tidy       # clean up dependencies
go mod download   # download dependencies
```

## Architecture

The project follows a standard Go CLI application structure:

- **Entry Point**: `main.go` - Uses Charmbracelet's fang for execution context
- **Command Structure**: `cmd/` directory contains Cobra command definitions
  - `cmd/root.go` - Defines the root command with configuration initialization
- **Configuration**: Supports multiple configuration sources:
  - Environment variables via `.env` file (using godotenv)
  - Config file (`.config.toml` in home directory or current directory)
  - Command-line flags (e.g., `--debug`, `--config`)
- **Logging**: Uses structured logging with slog and slogcolor for colored output
  - Debug mode can be enabled via `--debug` flag or config

## Key Dependencies

- **cobra**: Command-line interface framework
- **viper**: Configuration management (env vars, config files, flags)
- **charmbracelet/fang**: Enhanced command execution with context
- **slogcolor**: Colored structured logging
- **godotenv**: .env file support

## Development Notes

- The root command is named "example" and should be renamed for your application
- Debug logging is available via the `--debug` flag
- Configuration precedence: flags > env vars > config file > defaults
- Logger instance is globally available as `cmd.Logger`