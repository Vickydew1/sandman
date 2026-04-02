package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version is set at build time via -ldflags.
var Version = "dev"

var rootCmd = &cobra.Command{
	Use:   "sandman",
	Short: "Sandman: Security scanner powered by Trivy, Opengrep, ZAP, and ClamAV",
	Long:  `Sandman is a unified security scanner that wraps Trivy, Opengrep, OWASP ZAP, and ClamAV into a single CLI. It supports container image scanning, secrets detection, SAST, IaC misconfiguration checks, OS/package vulnerability scanning, malware detection, and DAST against live targets.`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the Sandman version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sandman %s\n", Version)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}