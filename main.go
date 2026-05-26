// crush is a terminal-based AI chat application.
// It is a fork of charmbracelet/crush with additional features and improvements.
package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/crush/internal/app"
	"github.com/charmbracelet/crush/internal/config"
	"github.com/spf13/cobra"
)

var (
	// Version is set at build time via ldflags.
	Version = "dev"
	// Commit is set at build time via ldflags.
	Commit = "none"
	// BuildDate is set at build time via ldflags.
	BuildDate = "unknown"
)

func main() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func rootCmd() *cobra.Command {
	var cfgFile string

	cmd := &cobra.Command{
		Use:   "crush",
		Short: "A terminal-based AI chat application",
		Long: `crush is a terminal-based AI chat application that lets you
converse with AI models directly from your terminal.

It supports multiple AI providers and offers a rich TUI experience
powered by Bubble Tea and Lip Gloss.`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load(cfgFile)
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			a, err := app.New(cfg)
			if err != nil {
				return fmt.Errorf("failed to initialize app: %w", err)
			}

			return a.Run()
		},
	}

	cmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default: $HOME/.config/crush/config.yaml)",
	)

	cmd.AddCommand(versionCmd())
	cmd.AddCommand(configCmd())

	return cmd
}

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("crush %s (commit: %s, built: %s)\n", Version, Commit, BuildDate)
		},
	}
}

func configCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage crush configuration",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "path",
		Short: "Print the path to the config file",
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := config.DefaultPath()
			if err != nil {
				return err
			}
			fmt.Println(path)
			return nil
		},
	})

	return cmd
}
