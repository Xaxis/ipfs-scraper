package server

import (
	"encoding/json"
	"github.com/Xaxis/ipfs-scraper/internal/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	DB *db.Database
}

func NewServer(db *db.Database) *Server {
	server := &Server{DB: db}
	return server
}

func (s *Server) Start(address string) {
	router := mux.NewRouter()
	router.HandleFunc("/tokens", s.handleTokens()).Methods("GET")
	router.HandleFunc("/tokens/{cid}", s.handleTokenByCID()).Methods("GET")
	log.Printf("Starting server on %s", address)
	log.Fatal(http.ListenAndServe(address, router))
}

// handleTokens handles the /tokens endpoint.
func (s *Server) handleTokens() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokens, err := s.DB.GetAllTokens()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(tokens)
	}
}

// handleTokenByCID handles the /tokens/{cid} endpoint.
func (s *Server) handleTokenByCID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cid := vars["cid"]
		token, err := s.DB.GetTokenByCID(cid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(token)
	}
}
