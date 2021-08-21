package handler

import "net/http"

// Handler turn http.Handler to pod.Handler
type Handler struct {
	http.Handler
}

// New Handler
func New(h http.Handler) *Handler {
	return &Handler{h}
}

// ServeHTTP implement pod.Handler
func (h *Handler) ServeHTTP(
	rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h.Handler.ServeHTTP(rw, r)
	next(rw, r)
}
