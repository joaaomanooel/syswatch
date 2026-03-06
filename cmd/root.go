package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "syswatch",
	Short:   "SysWatch is a CLI system monitor",
	Long:    `SysWatch provides real-time monitoring of CPU, memory, disk and process count.`,
	Version: version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of syswatch",
	Long:  `All software has versions. This is syswatch's.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("syswatch version", version)
	},
}
