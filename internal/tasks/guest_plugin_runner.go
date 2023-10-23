package tasks

import (
	"fmt"
	"kvm-agent/internal/config"
	"kvm-agent/internal/log"
	"kvm-agent/internal/runner/router"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server interface {
	ListenAndServe() error
}

func InitServer(address string, r *gin.Engine) Server {
	server := &http.Server{
		Addr:           address,
		Handler:        r,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server
}

func StartGuestPluginTask(config config.Server) {
	serverAddress := fmt.Sprintf("%s:%d", config.IP, config.Port)

	r := router.InitRouter()
	s := InitServer(serverAddress, r)
	log.Error("server start error: %v", s.ListenAndServe().Error())
}
