package metrics

import (
	"errors"
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

func TestMetricsSuite(t *testing.T) {
	suite.Run(t, new(MetricsTestSuite))
}
