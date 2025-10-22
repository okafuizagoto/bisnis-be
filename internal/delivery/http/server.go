package http

import (
	"net/http"

	"bisnis-be/pkg/grace"

	"github.com/rs/cors"
)

// GoldGymHandler ...
type GoldGymHandler interface {
	GetGoldGym(w http.ResponseWriter, r *http.Request)
	InsertGoldGym(w http.ResponseWriter, r *http.Request)
	DeleteGoldGym(w http.ResponseWriter, r *http.Request)
	UpdateGoldGym(w http.ResponseWriter, r *http.Request)
}

// AuthHandler ...
type AuthHandler interface {
	LoginUser(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	Goldgym GoldGymHandler
	Auth    AuthHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
