package web

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

type Server struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *Server {
	return &Server{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (w *Server) AddHandler(path string, handler http.HandlerFunc) {
	w.Handlers[path] = handler
}

// Start loop through all handlers and add them to the router
// register middleware logger and start the server
func (w *Server) Start() {
	w.Router.Use(middleware.Logger)

	for path, handler := range w.Handlers {
		w.Router.Handle(path, handler)
	}

	http.ListenAndServe(w.WebServerPort, w.Router)
}
