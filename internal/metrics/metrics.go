package metrics

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

const (
	clearScreen    = "\033[2J"
	moveCursorHome = "\033[H"
	hideCursor     = "\033[?25l"
	showCursor     = "\033[?25h"

	moveToCPUValue     = "\033[3;16H"
	moveToMemoryValue  = "\033[4;16H"
	moveToDiskValue    = "\033[5;16H"
	moveToProcessValue = "\033[6;16H"
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

	var cpuValue float64
	if len(cpuPct) > 0 {
		cpuValue = cpuPct[0]
	}

	return &Data{
		CPUPercent:   cpuValue,
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

type StreamPrinter struct {
	out io.Writer
}

func NewStreamPrinter(out io.Writer) *StreamPrinter {
	if out == nil {
		out = os.Stdout
	}
	return &StreamPrinter{out: out}
}

func (sp *StreamPrinter) Start() {
	// Combine multiple writes into a single buffer
	_, _ = fmt.Fprintf(sp.out, "%s%s%sSysWatch - Real-time System Metrics\nPress Ctrl+C to exit\nCPU Usage:      0.00%%\nMemory Usage:   0.00%%\nDisk Usage:     0.00%%\nProcesses:      0\n",
		hideCursor, clearScreen, moveCursorHome)
}

func (sp *StreamPrinter) Update(d *Data) {
	// Use a single formatted string for all updates
	_, _ = fmt.Fprintf(sp.out, "%s%.2f%%%s%.2f%%%s%.2f%%%s%d   ",
		moveToCPUValue, d.CPUPercent,
		moveToMemoryValue, d.MemPercent,
		moveToDiskValue, d.DiskPercent,
		moveToProcessValue, d.ProcessCount)
}

// Stop cleans up the terminal state
func (sp *StreamPrinter) Stop() {
	_, _ = fmt.Fprintf(sp.out, "%s\n\nMonitoring stopped.\n", showCursor)
}
