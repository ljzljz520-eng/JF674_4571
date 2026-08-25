package api

import (
	"encoding/json"
	"galleryline/service"
	"net/http"
)

type Server struct {
	dir   *service.Directory
	calls *service.CallManager
}

func NewServer(d *service.Directory, c *service.CallManager) *Server {
	return &Server{dir: d, calls: c}
}
func (s *Server) Handler() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNoContent) })
	m.HandleFunc("/extensions", s.extensions)
	m.HandleFunc("/records", s.records)
	return m
}
func (s *Server) extensions(w http.ResponseWriter, r *http.Request) {
	v, e := s.dir.List()
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(v)
}
func (s *Server) records(w http.ResponseWriter, r *http.Request) {
	v, e := s.calls.History()
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(v)
}
