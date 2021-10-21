package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("\\") // If you're in Unix change this "\\" for "/"

	fmt.Println("Bilgisayar Adı:", hostStat.Hostname)
	fmt.Println("İşletim Sistemi:", hostStat.Platform)
	fmt.Println("İşlemci:", cpuStat[0].ModelName)
	fmt.Println("RAM:", vmStat.Total/1024/1024)
	fmt.Println("Disk:", diskStat.Total/1024/1024)
}
