package cmd

import (
	"bytes"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MonitorTestSuite struct {
	suite.Suite
	buffer         *bytes.Buffer
	originalTicker func(d time.Duration) *time.Ticker
	stop           chan struct{}
}

func (s *MonitorTestSuite) SetupTest() {
	rootCmd := &cobra.Command{
		Use:   "syswatch",
		Short: "syswatch monitor test suite",
	}

	s.buffer = &bytes.Buffer{}

	s.stop = make(chan struct{})
	s.originalTicker = tickerFactory

	tickerFactory = func(d time.Duration) *time.Ticker {
		t := time.NewTicker(d)
		go func() {
			<-t.C
			close(s.stop)
		}()
		return t
	}

	rootCmd.ResetCommands()
	rootCmd.AddCommand(MonitorCmd())
	rootCmd.SetOut(s.buffer)
	rootCmd.SetErr(s.buffer)
}

func (s *MonitorTestSuite) TearDownTest() {
	tickerFactory = s.originalTicker
}

func (s *MonitorTestSuite) TestMonitorCommand_OutputsMetrics() {
	rootCmd.SetArgs([]string{"monitor"})

	var bufOut bytes.Buffer
	rootCmd.SetOut(&bufOut)

	go func() {
		err := rootCmd.Execute()

		assert.NoError(s.T(), err)

		out := s.buffer.String()
		assert.Contains(s.T(), out, "CPU Usage:", "deve exibir uso de CPU")
		assert.Contains(s.T(), out, "Memory Usage:", "deve exibir uso de memória")
		assert.Contains(s.T(), out, "Disk Usage:", "deve exibir uso de disco")
		assert.Contains(s.T(), out, "Processes:", "deve exibir contagem de processors")
	}()
	<-s.stop
}

func (s *MonitorTestSuite) TestMonitorCommand_WithCustomInterval() {
	rootCmd.SetArgs([]string{"monitor", "-i", "2s"})

	tickerFactory = func(d time.Duration) *time.Ticker {
		assert.Equal(s.T(), time.Second*2, d, "interval should be 2s")
		t := time.NewTicker(d)
		go func() {
			<-t.C
			close(s.stop)
		}()
		return t
	}

	go func() {
		err := rootCmd.Execute()
		assert.NoError(s.T(), err)
	}()
	<-s.stop
}

func TestMonitorSuit(t *testing.T) {
	suite.Run(t, new(MonitorTestSuite))
}
