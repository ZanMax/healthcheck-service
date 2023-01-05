package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	http "net/http"
	"time"
)

type HealthCheck struct {
	Status string `json:"status"`
	Date   string `json:"date"`
}

type HealthCheckExtended struct {
	Status string          `json:"status"`
	Date   string          `json:"date"`
	Host   hostInformation `json:"host"`
	Stats  serverStats     `json:"stats"`
}

type serverStats struct {
	CpuUsage  float64 `json:"cpu_usage"`
	MemUsage  float64 `json:"mem_usage"`
	DiskUsage float64 `json:"disk_usage"`
}

type hostInformation struct {
	Hostname        string `json:"hostname"`
	OS              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platformVersion"`
	KernelVersion   string `json:"kernelVersion"`
	KernelArch      string `json:"kernelArch"`
	Uptime          uint64 `json:"uptime"`
	BootTime        uint64 `json:"bootTime"`
	Procs           uint64 `json:"procs"`
}

func getHealthCheck(c *gin.Context) {
	currentTime := time.Now()
	hc := HealthCheck{Status: "OK", Date: currentTime.Format("2006-01-02 15:04:05.000000000")}
	c.JSON(http.StatusOK, hc)
}

func getHealthCheckExtended(c *gin.Context) {
	currentTime := time.Now()
	var cp, _ = cpu.Percent(time.Second, true)
	var v, _ = mem.VirtualMemory()
	var d, _ = disk.Usage("/")
	hostInfo, _ := host.Info()
	additionalHostInfo := hostInformation{
		Hostname:        hostInfo.Hostname,
		OS:              hostInfo.OS,
		Platform:        hostInfo.Platform,
		PlatformVersion: hostInfo.PlatformVersion,
		KernelVersion:   hostInfo.KernelVersion,
		KernelArch:      hostInfo.KernelArch,
		Uptime:          hostInfo.Uptime,
		BootTime:        hostInfo.BootTime,
		Procs:           hostInfo.Procs,
	}

	var hcExt = HealthCheckExtended{Status: "OK", Date: currentTime.Format("2006-01-02 15:04:05.000000000"), Host: additionalHostInfo, Stats: serverStats{CpuUsage: cp[0], MemUsage: v.UsedPercent, DiskUsage: d.UsedPercent}}
	c.JSON(http.StatusOK, hcExt)
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"service": "healthcheck",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", index)
	router.GET("/healthcheck", getHealthCheck)
	router.GET("/healthcheck/ext", getHealthCheckExtended)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
