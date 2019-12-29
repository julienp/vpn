package server

import (
	"encoding/json"
	"log"
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
			log.Printf("Could not refresh status: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			var location *string
			if s.vpn.Location != nil {
				location = &s.vpn.Location.Alias
			}
			s.respond(w, r, map[string]interface{}{"status": s.vpn.Status, "location": location}, http.StatusOK)
		}
	}
}

func (s *server) handleListLocations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		locations, err := s.vpn.ListLocations()
		if err != nil {
			log.Printf("Could not get locations: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			s.respond(w, r, map[string]interface{}{"locations": locations}, http.StatusOK)
		}
	}
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	if data != nil {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf("Could not encode json=%v err=%s", data, err)
		}
		if status != http.StatusOK {
			w.WriteHeader(status)
		}
	} else {
		w.WriteHeader(status)
	}
}

func NewServer() http.Handler {
	srv := &server{
		vpn:    controller.NewController(),
		router: http.NewServeMux(),
	}
	srv.router.HandleFunc("/status", srv.handleStatus())
	srv.router.HandleFunc("/locations", srv.handleListLocations())
	return srv
}
