package handlers

import (
	"fmt"
	"kvm-agent/internal/log"
	"kvm-agent/internal/runner/models/request"
	"kvm-agent/internal/runner/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunPingPlug(c *gin.Context) {
	var pingPlug request.PingPlug
	err := c.ShouldBindJSON(&pingPlug)
	if err != nil {
		fmt.Println("c.ShouldBindJSON error:", err)
		log.Errorf("RunPingPlug", "c.ShouldBindJSON error: %v", err)

		return
	}

	s := service.NewDefaultService(c)
	result := s.HandlePing(pingPlug)

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    result,
	})
}

func RunFIOPlug(c *gin.Context) {
	var fioPlug request.FIOPlug
	err := c.ShouldBindJSON(&fioPlug)
	if err != nil {
		fmt.Println("c.ShouldBindJSON error:", err)
		log.Errorf("RunFIOPlug", "c.ShouldBindJSON error: %v", err)

		return
	}

	s := service.NewDefaultService(c)
	result := s.HandleFIO(fioPlug)

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    result,
	})
}

func RunDBTestPlug(c *gin.Context) {
	var dbTestPPlug request.DBTestPlug
	err := c.ShouldBindJSON(&dbTestPPlug)
	if err != nil {
		fmt.Println("c.ShouldBindJSON error:", err)
		log.Errorf("RunDBTestPlug", "c.ShouldBindJSON error: %v", err)

		return
	}

	s := service.NewDefaultService(c)
	result := s.HandleDBTest(dbTestPPlug)

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    result,
	})
}

func RunPTP4LPlug(c *gin.Context) {
	var ptp4lPlug request.PTP4LPlug
	err := c.ShouldBindJSON(&ptp4lPlug)
	if err != nil {
		fmt.Println("c.ShouldBindJSON error:", err)
		log.Errorf("RunPTP4LPlug", "c.ShouldBindJSON error: %v", err)

		return
	}

	s := service.NewDefaultService(c)
	result := s.HandlePTP4L(ptp4lPlug)

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    result,
	})
}
