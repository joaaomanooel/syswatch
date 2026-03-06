package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joaaomanooel/syswatch/internal/metrics"
	"github.com/spf13/cobra"
)

var interval time.Duration

var tickerFactory = func(d time.Duration) *time.Ticker {
	return time.NewTicker(d)
}

func handleMonitor(cmd *cobra.Command, args []string) error {
	ticker := tickerFactory(interval)
	defer ticker.Stop()

	// Set up signal handling for clean exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	provider := &metrics.RealProvider{}

	printer := metrics.NewStreamPrinter(os.Stdout)
	printer.Start()
	defer printer.Stop()

	for {
		select {
		case <-ticker.C:
			data, err := metrics.CollectAll(provider)
			if err != nil {
				return err
			}
			printer.Update(data)
		case <-sigChan:
			return nil
		}
	}
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
