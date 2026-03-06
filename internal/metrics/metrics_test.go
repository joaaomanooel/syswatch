package metrics

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type mockProvider struct {
	cpuData []float64
	vmData  *mem.VirtualMemoryStat
	duData  *disk.UsageStat
	procs   []*process.Process
	err     error
}

func (m mockProvider) CPUPercent(i time.Duration, p bool) ([]float64, error) {
	return m.cpuData, m.err
}
func (m mockProvider) VirtualMemory() (*mem.VirtualMemoryStat, error) {
	return m.vmData, m.err
}
func (m mockProvider) DiskUsage(path string) (*disk.UsageStat, error) {
	return m.duData, m.err
}
func (m mockProvider) Processes() ([]*process.Process, error) {
	return m.procs, m.err
}

type MetricsTestSuite struct {
	suite.Suite
}

func (t *MetricsTestSuite) TestCollectAll_Success() {
	sampleCPU := []float64{12.34}
	sampleVM := &mem.VirtualMemoryStat{UsedPercent: 56.78}
	sampleDU := &disk.UsageStat{UsedPercent: 90.12}
	sampleProcs := make([]*process.Process, 5)

	provider := mockProvider{
		cpuData: sampleCPU,
		vmData:  sampleVM,
		duData:  sampleDU,
		procs:   sampleProcs,
		err:     nil,
	}

	result, err := CollectAll(provider)

	assert.NoError(t.T(), err, "should not return an error on success")
	assert.Equal(t.T(), 12.34, result.CPUPercent, "CPUPercent should match")
	assert.Equal(t.T(), 56.78, result.MemPercent, "MemPercent should match")
	assert.Equal(t.T(), 90.12, result.DiskPercent, "DiskPercent should match")
	assert.Equal(t.T(), 5, result.ProcessCount, "ProcessCount should match")
}

func (t *MetricsTestSuite) TestCollectAll_Error() {
	provider := mockProvider{err: errors.New("fail")}
	result, err := CollectAll(provider)

	assert.Error(t.T(), err, "should return an error when provider fails")
	assert.Nil(t.T(), result, "result should be nil on error")
}

func (t *MetricsTestSuite) TestCollectAll_EmptyCPU() {
	provider := mockProvider{
		cpuData: []float64{},
		vmData:  &mem.VirtualMemoryStat{UsedPercent: 50.0},
		duData:  &disk.UsageStat{UsedPercent: 50.0},
		procs:   []*process.Process{},
		err:     nil,
	}

	result, err := CollectAll(provider)

	assert.NoError(t.T(), err)
	assert.NotNil(t.T(), result)
	assert.Equal(t.T(), 0.0, result.CPUPercent, "CPUPercent should be 0 for empty slice")
}

func (t *MetricsTestSuite) TestPrint() {
	data := &Data{
		CPUPercent:   25.5,
		MemPercent:   60.2,
		DiskPercent:  75.8,
		ProcessCount: 100,
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Print(data)

	if err := w.Close(); err != nil {
		assert.NoError(t.T(), err)
	}

	os.Stdout = old

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		assert.NoError(t.T(), err)
	}

	output := buf.String()

	assert.Contains(t.T(), output, "25.50%")
	assert.Contains(t.T(), output, "60.20%")
	assert.Contains(t.T(), output, "75.80%")
	assert.Contains(t.T(), output, "100")
}

// Remove any TestStreamPrinter method from the suite
// The standalone TestStreamPrinter function will handle this testing

func TestMetricsSuite(t *testing.T) {
	suite.Run(t, new(MetricsTestSuite))
}

func TestRealProvider(t *testing.T) {
	t.Run("RealProvider can collect CPUPercent", func(t *testing.T) {
		provider := &RealProvider{}
		cpu, err := provider.CPUPercent(0, false)
		assert.NoError(t, err)
		assert.NotNil(t, cpu)
	})

	t.Run("RealProvider can collect VirtualMemory", func(t *testing.T) {
		provider := &RealProvider{}
		vm, err := provider.VirtualMemory()
		assert.NoError(t, err)
		assert.NotNil(t, vm)
	})

	t.Run("RealProvider can collect DiskUsage", func(t *testing.T) {
		provider := &RealProvider{}
		du, err := provider.DiskUsage("/")
		assert.NoError(t, err)
		assert.NotNil(t, du)
	})

	t.Run("RealProvider can collect Processes", func(t *testing.T) {
		provider := &RealProvider{}
		procs, err := provider.Processes()
		assert.NoError(t, err)
		assert.NotNil(t, procs)
	})
}

// Keep the standalone TestStreamPrinter function as is
func TestStreamPrinter(t *testing.T) {
	t.Run("NewStreamPrinter with nil writer", func(t *testing.T) {
		sp := NewStreamPrinter(nil)
		assert.NotNil(t, sp)
		assert.Equal(t, os.Stdout, sp.out)
	})

	t.Run("NewStreamPrinter with custom writer", func(t *testing.T) {
		buf := &bytes.Buffer{}
		sp := NewStreamPrinter(buf)
		assert.NotNil(t, sp)
		assert.Equal(t, buf, sp.out)
	})

	t.Run("Start method writes expected output", func(t *testing.T) {
		buf := &bytes.Buffer{}
		sp := NewStreamPrinter(buf)
		sp.Start()

		expected := fmt.Sprintf("%s%s%sSysWatch - Real-time System Metrics\nPress Ctrl+C to exit\nCPU Usage:      0.00%%\nMemory Usage:   0.00%%\nDisk Usage:     0.00%%\nProcesses:      0\n",
			hideCursor, clearScreen, moveCursorHome)
		assert.Equal(t, expected, buf.String())
	})

	t.Run("Update method writes expected output", func(t *testing.T) {
		buf := &bytes.Buffer{}
		sp := NewStreamPrinter(buf)

		data := &Data{
			CPUPercent:   10.5,
			MemPercent:   45.2,
			DiskPercent:  75.8,
			ProcessCount: 123,
		}

		sp.Update(data)

		expected := fmt.Sprintf("%s%.2f%%%s%.2f%%%s%.2f%%%s%d   ",
			moveToCPUValue, data.CPUPercent,
			moveToMemoryValue, data.MemPercent,
			moveToDiskValue, data.DiskPercent,
			moveToProcessValue, data.ProcessCount)
		assert.Equal(t, expected, buf.String())
	})

	t.Run("Stop method writes expected output", func(t *testing.T) {
		buf := &bytes.Buffer{}
		sp := NewStreamPrinter(buf)

		sp.Stop()

		expected := fmt.Sprintf("%s\n\nMonitoring stopped.\n", showCursor)
		assert.Equal(t, expected, buf.String())
	})

	t.Run("Full lifecycle test", func(t *testing.T) {
		buf := &bytes.Buffer{}
		sp := NewStreamPrinter(buf)

		sp.Start()
		buf.Reset() // Clear buffer after Start

		data := &Data{
			CPUPercent:   10.5,
			MemPercent:   45.2,
			DiskPercent:  75.8,
			ProcessCount: 123,
		}

		sp.Update(data)
		sp.Stop()

		// Check that the buffer contains both Update and Stop outputs
		assert.Contains(t, buf.String(), "10.50%")
		assert.Contains(t, buf.String(), "45.20%")
		assert.Contains(t, buf.String(), "75.80%")
		assert.Contains(t, buf.String(), "123")
		assert.Contains(t, buf.String(), "Monitoring stopped")
	})
}
