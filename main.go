package main

import (
	"fmt"
	"net/http"
	"one-siam-restaurant/configs"
	"one-siam-restaurant/internal/restaurant"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	//Load configs
	cfg, err := configs.LoadConfig()
	if err != nil {
		fmt.Println("Error loading configs")
	}

	r := gin.Default()
	r.GET("/health-check", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": time.Now().Format(time.RFC3339),
		})
	})
	restaurant.Initialize(r, cfg)
	go r.Run()

	<-sig
}
