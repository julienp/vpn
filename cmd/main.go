package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienp/vpn/controller"
)

type server struct {
	vpn    *controller.Controller
	router *http.ServeMux
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) handleStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := s.vpn.RefreshStatus()
		if err != nil {
			w.WriteHeader(502)
		} else {
			w.WriteHeader(200)
			fmt.Fprintf(w, s.vpn.Status.String())
		}
	}
}

func newServer() *server {
	srv := &server{
		vpn:    controller.NewController(),
		router: http.NewServeMux(),
	}
	srv.router.HandleFunc("/status", srv.handleStatus())
	return srv
}

func main() {
	addr := "0.0.0.0:9999"
	srv := &http.Server{Addr: addr, Handler: newServer()}
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
