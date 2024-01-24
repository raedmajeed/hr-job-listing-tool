package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/hr-job-tool/pkg/di"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sign := make(chan os.Signal)
	signal.Notify(sign, os.Interrupt, syscall.SIGINT)

	r := gin.Default()
	server := di.Init(r)
	go server.StartServer()

	<-sign
	log.Println("program interrupted, gracefully shutting down server after 2 seconds")
	time.Sleep(time.Second * 2)
}
