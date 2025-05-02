package metrics

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

type Data struct {
	CPUPercent   float64
	MemPercent   float64
	DiskPercent  float64
	ProcessCount int
}

func CollectAll() (*Data, error) {
	pct, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	du, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	return &Data{
		CPUPercent:   pct[0],
		MemPercent:   vm.UsedPercent,
		DiskPercent:  du.UsedPercent,
		ProcessCount: len(procs),
	}, nil
}
