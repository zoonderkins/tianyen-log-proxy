package server

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/t9963/log-proxy-server/pkg/config"
	"gitlab.com/t9963/log-proxy-server/pkg/logger"
)

type LogRequest struct {
	Level           string            `json:"level"`
	Message         string            `json:"message"`
	Logger          string            `json:"logger"`
	Service         string            `json:"service"`
	Application     string            `json:"application"`
	UserAgent       string            `json:"user_agent"`
	LocationURLPath string            `json:"locationurlPath"`
	UserId          string            `json:"userId"`
	AdditionalInfo  map[string]string `json:"additionalInfo"`
}

func Start() {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/log", func(c *gin.Context) {
		var logReq LogRequest
		if err := c.ShouldBindJSON(&logReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		additionalInfo := map[string]string{
			"logger":          logReq.Logger,
			"service":         logReq.Service,
			"application":     logReq.Application,
			"logLevel":        logReq.Level,
			"user_agent":      logReq.UserAgent,
			"locationurlPath": logReq.LocationURLPath,
			"userId":          logReq.UserId,
		}

		for key, value := range logReq.AdditionalInfo {
			additionalInfo[key] = value
		}

		err := logger.SendLog(logReq.Level, logReq.Message, additionalInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Log sent successfully"})
	})

	// Print the environment variables
	config.PrintEnv()

	port := config.ServerPort
	fmt.Printf("Starting server on port %s\n", port)
	r.Run(":" + port)
}
