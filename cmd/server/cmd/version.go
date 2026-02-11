package cmd

import (
	"ddd/internal/version"
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		info := version.Get()
		fmt.Printf("Version: %s\nGitHash: %s\nBuildTime: %s\n",
			info.Version,
			info.GitHash,
			info.BuildTime,
		)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
