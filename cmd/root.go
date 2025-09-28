// Package cmd implements the command-line interface using Cobra and Viper.
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ErrRootCmd indicates a failure in root command execution.
var ErrRootCmd = errors.New("failed to run example command")

// NewRootCmd creates the root command with configured flags and pre-run hooks.
func NewRootCmd(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "An example application.",
		Long:  "An example application, it doesn't do anything.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlag("debug", cmd.PersistentFlags().Lookup("debug")); err != nil {
				return fmt.Errorf("%w: %s", ErrRootCmd, err)
			}

			return nil
		},
	}

	var configFile string
	var debug bool

	cmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.config.toml)")
	cmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug mode")

	// cmd.AddCommand(NewChildCommand())

	return cmd
}

// Execute initializes the logger, configuration, and runs the root command.
// It returns the command instance regardless of execution success.
func Execute() *cobra.Command {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		Level:           log.WarnLevel,
	})

	cobra.OnInitialize(func() {
		initConfig(logger)
	})

	cmd := NewRootCmd(logger)

	if err := cmd.Execute(); err != nil {
		logger.Error(ErrRootCmd.Error(), "error", err)

		return nil
	}

	return cmd
}

// initConfig loads configuration from environment variables, config files, and flags.
// Configuration precedence: flags > env > config file > defaults.
func initConfig(logger *log.Logger) {
	if err := godotenv.Load(); err != nil {
		logger.Debug(".env file not found, using environment variables")
	} else {
		logger.Debug(".env file loaded successfully")
	}

	configFile := viper.GetString("config")

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".config")
		viper.SetConfigType("toml")
	}

	viper.AutomaticEnv()
	if err := viper.BindEnv("debug", "DEBUG"); err != nil {
		logger.Error(ErrRootCmd.Error(), "error", err)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Debug("config file not found")
		} else {
			logger.Error("error loading config file", "error", err)
		}
	} else {
		logger.Debug("using config file", "file", viper.ConfigFileUsed())
	}

	if viper.GetBool("debug") {
		logger.SetLevel(log.DebugLevel)
	}
}
