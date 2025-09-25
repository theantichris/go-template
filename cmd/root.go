package cmd

import (
	"log/slog"
	"os"
	"time"

	"github.com/MatusOllah/slogcolor"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
	Debug      bool
	Logger     *slog.Logger
	EnvVar     string
)

// RootCmd is the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "example",
	Short: "An example application.",
	Long:  "An example application, it doesn't do anything.",
}

// Execute adds initialization.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		Logger.Error("error running command", "component", "cmd.RootCmd")
		os.Exit(1)
	}
}

// init sets and binds flags.
func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.config.toml)")
	RootCmd.PersistentFlags().BoolVar(&Debug, "debug", false, "enable debug mode")

	viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug"))
}

// initConfig loads env variables and the config file.
func initConfig() {
	initLogger()

	if err := godotenv.Load(); err != nil {
		Logger.Debug(".env file not found, using environment variables", "component", "cmd.RootCmd")
	} else {
		Logger.Debug(".env file loaded successfully", "component", "cmd.RootCmd")
	}

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
	viper.BindEnv("envVar", "ENV_VAR")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			Logger.Debug("config file not found", "component", "cmd.RootCmd")
		} else {
			Logger.Error("error loading config file", "error", err)
		}
	} else {
		Logger.Debug("using config file", "file", viper.ConfigFileUsed(), "component", "cmd.RootCmd")
	}

	if viper.GetBool("debug") {
		Debug = true
		initLogger()
	}
}

// initLogger initializes the logger.
func initLogger() {
	logLevel := slog.LevelWarn

	if Debug {
		logLevel = slog.LevelDebug
	}

	Logger = slog.New(slogcolor.NewHandler(os.Stderr, &slogcolor.Options{
		Level:      logLevel,
		TimeFormat: time.RFC3339,
	}))

	slog.SetDefault(Logger)
}
