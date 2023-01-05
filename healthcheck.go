package main

import (
	"github.com/gin-gonic/gin"
	http "net/http"
	"time"
)

type HealthCheck struct {
	Status string `json:"status"`
	Date   string `json:"date"`
}

type HealthCheckExtended struct {
	Status string      `json:"status"`
	Date   string      `json:"date"`
	Stats  serverStats `json:"stats"`
}

type serverStats struct {
	CpuUsage  float64 `json:"cpu_usage"`
	MemUsage  float64 `json:"mem_usage"`
	DiskUsage float64 `json:"disk_usage"`
}

func getHealthCheck(c *gin.Context) {
	currentTime := time.Now()
	hc := HealthCheck{Status: "OK", Date: currentTime.Format("2006-01-02 15:04:05.000000000")}
	c.JSON(http.StatusOK, hc)
}

func getHealthCheckExtended(c *gin.Context) {
	currentTime := time.Now()
	var hcExt = HealthCheckExtended{Status: "OK", Date: currentTime.Format("2006-01-02 15:04:05.000000000"), Stats: serverStats{CpuUsage: 0.5, MemUsage: 0.5, DiskUsage: 0.5}}
	c.IndentedJSON(http.StatusOK, hcExt)
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
