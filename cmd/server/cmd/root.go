package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	configFile string
)

var rootCmd = &cobra.Command{
	Use:   "ddd",
	Short: "ddd service",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// 全局参数
	rootCmd.PersistentFlags().StringVar(
		&configFile,
		"config",
		"",
		"config file path (optional)",
	)
}
