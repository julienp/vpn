package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienp/vpn/status"
)

type expressVPN struct {
	command   string
	extraArgs []string
	status    status.VPNStatus
}

func (e *expressVPN) RefreshStatus() error {
	args := append(e.extraArgs, "status")
	cmd := exec.Command(e.command, args...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	e.status = status.ParseStatus(string(stdoutStderr))
	return nil
}

type server struct {
	vpn    *expressVPN
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
			fmt.Fprintf(w, s.vpn.status.String())
		}
	}
}

func newServer() *server {
	srv := &server{
		vpn: &expressVPN{
			command:   "sh",
			extraArgs: []string{"-c", "sleep 2 && echo 'Connected to lala'"},
		},
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
