package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MetricsTestSuite struct {
	suite.Suite
}

func (t *MetricsTestSuite) TestCollectAll_ReturnsNonNilData() {
	data, err := CollectAll()

	assert.NoError(t.T(), err, "CollectAll should not return an error")
	assert.NotNil(t.T(), data, "Data should not be nil")
}

func (t *MetricsTestSuite) TestCollectAll_MetricsWithinRange() {
	data, err := CollectAll()

	assert.NoError(t.T(), err, "CollectAll should not return an error")
	assert.GreaterOrEqual(t.T(), data.CPUPercent, 0.0, "CPU percent should be >= 0")
	assert.LessOrEqual(t.T(), data.CPUPercent, 100.0, "CPU percent should be <= 100")

	assert.GreaterOrEqual(t.T(), data.MemPercent, 0.0, "Memory percent should be >= 0")
	assert.LessOrEqual(t.T(), data.MemPercent, 100.0, "Memory percent should be <= 100")

	assert.GreaterOrEqual(t.T(), data.DiskPercent, 0.0, "Disk percent should be >= 0")
	assert.LessOrEqual(t.T(), data.DiskPercent, 100.0, "Disk percent should be <= 100")

	assert.Greater(t.T(), data.ProcessCount, 0, "Process count should be > 0")
}

func TestMetricsSuite(t *testing.T) {
	suite.Run(t, new(MetricsTestSuite))
}
