package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
