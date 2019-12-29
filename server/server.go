package server

import (
	"fmt"
	"net/http"

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

func NewServer() http.Handler {
	srv := &server{
		vpn:    controller.NewController(),
		router: http.NewServeMux(),
	}
	srv.router.HandleFunc("/status", srv.handleStatus())
	return srv
}
