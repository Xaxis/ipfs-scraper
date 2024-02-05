package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Xaxis/ipfs-scraper/internal/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	DB *db.Database
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type TokensResponse struct {
	Tokens []db.IPFSMetadata `json:"tokens"`
}

func NewServer(db *db.Database) *Server {
	server := &Server{DB: db}
	return server
}

func (s *Server) Start(address string) {
	router := mux.NewRouter()
	router.HandleFunc("/tokens", s.handleTokens()).Methods("GET")
	router.HandleFunc("/tokens/{cid}", s.handleTokenByCID()).Methods("GET")

	router.Use(loggingMiddleware)

	srv := &http.Server{
		Addr:         address,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Starting server on %s", address)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Shutting down gracefully, press Ctrl+C again to force")
}

// loggingMiddleware logs all incoming requests.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// handleTokens handles the /tokens endpoint.
func (s *Server) handleTokens() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokens, err := s.DB.GetAllTokens()
		if err != nil {
			log.Printf("Error fetching tokens: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Internal Server Error"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TokensResponse{Tokens: tokens})
	}
}

// handleTokenByCID handles the /tokens/{cid} endpoint.
func (s *Server) handleTokenByCID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cid := vars["cid"]
		token, err := s.DB.GetTokenByCID(cid)
		if err != nil {
			log.Printf("Error fetching token by CID %s: %v", cid, err)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResponse{Error: fmt.Sprintf("Token not found for CID: %s", cid)})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(token)
	}
}
