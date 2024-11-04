package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Armunz/golang-simple-cron/internal/config"
	"github.com/robfig/cron/v3"
)

func main() {
	cfg := config.New()

	c := cron.New()
	c.AddFunc(cfg.CronTime, func() {
		log.Println("Hello World")
	})

	log.Println("cron job started")
	go func() {
		c.Start()
	}()

	// termination signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	log.Println("termination signal detected...")

	log.Println("closing cron job")
	cronCtx := c.Stop()
	<-cronCtx.Done()
}
