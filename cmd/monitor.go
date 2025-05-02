package cmd

import (
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/joaaomanooel/syswatch/internal/metrics"
	"github.com/spf13/cobra"
)

var interval time.Duration

var tickerFactory = func(d time.Duration) *time.Ticker {
	return time.NewTicker(d)
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func handleMonitor(cmd *cobra.Command, args []string) error {
	ticker := tickerFactory(interval)
	defer ticker.Stop()

	for range ticker.C {
		clearScreen()

		data, err := metrics.CollectAll(&metrics.RealProvider{})
		if err != nil {
			return err
		}

		metrics.Print(data)
	}
	return nil
}

func MonitorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor system resources in real time",
		Long:  `Displays CPU, memory, disk and process count metrics at configurable intervals.`,
		RunE:  handleMonitor,
	}

	cmd.Flags().DurationVarP(&interval, "interval", "i", 2*time.Second, "Update interval (e.g. 500ms, 1s)")
	return cmd
}

func init() {
	rootCmd.AddCommand(MonitorCmd())
}
