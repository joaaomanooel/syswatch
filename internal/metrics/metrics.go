package metrics

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

type MetricsProvider interface {
	CPUPercent(interval time.Duration, percpu bool) ([]float64, error)
	VirtualMemory() (*mem.VirtualMemoryStat, error)
	DiskUsage(path string) (*disk.UsageStat, error)
	Processes() ([]*process.Process, error)
}

type RealProvider struct{}

func (rp *RealProvider) CPUPercent(interval time.Duration, percpu bool) ([]float64, error) {
	return cpu.Percent(interval, percpu)
}
func (rp *RealProvider) VirtualMemory() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}
func (rp *RealProvider) DiskUsage(path string) (*disk.UsageStat, error) {
	return disk.Usage(path)
}
func (rp *RealProvider) Processes() ([]*process.Process, error) {
	return process.Processes()
}

type Data struct {
	CPUPercent   float64
	MemPercent   float64
	DiskPercent  float64
	ProcessCount int
}

func CollectAll(provider MetricsProvider) (*Data, error) {
	cpuPct, err := provider.CPUPercent(0, false)
	if err != nil {
		return nil, err
	}

	vm, err := provider.VirtualMemory()
	if err != nil {
		return nil, err
	}

	du, err := provider.DiskUsage("/")
	if err != nil {
		return nil, err
	}

	procs, err := provider.Processes()
	if err != nil {
		return nil, err
	}

	return &Data{
		CPUPercent:   cpuPct[0],
		MemPercent:   vm.UsedPercent,
		DiskPercent:  du.UsedPercent,
		ProcessCount: len(procs),
	}, nil
}

func Print(d *Data) {
	fmt.Printf("CPU Usage:      %.2f%%\n", d.CPUPercent)
	fmt.Printf("Memory Usage:   %.2f%%\n", d.MemPercent)
	fmt.Printf("Disk Usage:     %.2f%%\n", d.DiskPercent)
	fmt.Printf("Processes:      %d\n\n", d.ProcessCount)
}
