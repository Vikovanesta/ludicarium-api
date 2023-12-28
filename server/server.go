package server

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/vikovanesta/ludicarium-api/igdb"
)

type Server struct {
	server     *http.Server
	router     *mux.Router
	igdbClient *igdb.Client
}

func NewServer() *Server {
	s := Server{
		server: &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  10 * time.Second,
		},
		router: mux.NewRouter().StrictSlash(true),
	}

	s.routes()

	s.igdbClient = igdb.NewClient(nil)
	s.server.Handler = s.router

	return &s
}

func (s *Server) Run(port string) error {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	s.server.Addr = port
	log.Printf("Server running on port %s", port)
	return s.server.ListenAndServe()
}

func healthCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := M{
			"status":  "ok",
			"message": "Server is running",
		}
		writeJson(w, http.StatusOK, response)
	})
}
