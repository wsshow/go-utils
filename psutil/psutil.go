package psutil

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func diskInfo() {
	fmt.Println("--------------------DiskInfo--------------------")
	//显示磁盘分区信息
	partitions, _ := disk.Partitions(true)
	var totalSize uint64 = 0
	for _, part := range partitions {
		//fmt.Printf("part:%v\n", part.String())
		usage, _ := disk.Usage(part.Mountpoint)
		partSize := usage.Total >> 30
		totalSize += partSize
		fmt.Printf("[%v 磁盘大小:%4vGB] 磁盘使用率:%5.2f%% 磁盘剩余空间:%4vGB\n", part.Device, partSize, usage.UsedPercent, usage.Free>>30)
	}
	fmt.Printf(">>>>>>>> 磁盘总大小: %vGB\n", totalSize)
}

func memInfo() {
	fmt.Println("--------------------MemInfo--------------------")
	v, _ := mem.VirtualMemory()
	fmt.Printf("内存总大小: %vGB, 剩余内存大小:%vGB, 内存使用率:%.2f%%\n", v.Total>>30, v.Free>>30, v.UsedPercent)
}

func WorkFlow() {
	diskInfo()
	memInfo()
}
