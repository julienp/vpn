package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienp/vpn/server"
)

func main() {
	addr := "0.0.0.0:9999"
	srv := &http.Server{Addr: addr, Handler: server.NewServer()}
	log.Printf("Starting server on %s", addr)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				log.Print(err)
			} else {
				log.Print(err)
			}
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Printf("Stopping ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
