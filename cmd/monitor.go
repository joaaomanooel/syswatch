package cmd

import (
	"fmt"
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
	cmd.Run()
}

func handleMonitor(cmd *cobra.Command, args []string) error {
	ticker := tickerFactory(interval)
	defer ticker.Stop()

	for range ticker.C {
		clearScreen()

		data, err := metrics.CollectAll()
		if err != nil {
			return err
		}

		// Single formatted print instead of multiple prints
		fmt.Printf("CPU Usage:    %.2f%%\n", data.CPUPercent)
		fmt.Printf("Memory Usage: %.2f%%\n", data.MemPercent)
		fmt.Printf("Disk Usage:   %.2f%%\n", data.DiskPercent)
		fmt.Printf("Processes:    %d\n\n", data.ProcessCount)
	}
	return nil
}

func MonitorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor system resources in real time",
		Long:  `Exibe uso de CPU, memória, disco e contagem de processos em tempo real.`,
		RunE:  handleMonitor,
	}

	cmd.Flags().DurationVarP(&interval, "interval", "i", 2*time.Second, "Update interval (e.g. 500ms, 1s)")
	return cmd
}

func init() {
	rootCmd.AddCommand(MonitorCmd())
}
