package web

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (w *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	w.Handlers[path] = handler
}

// Start loop through all handlers and add them to the router
// register middleware logger and start the server
func (w *WebServer) Start() {
	w.Router.Use(middleware.Logger)

	for path, handler := range w.Handlers {
		w.Router.Handle(path, handler)
	}

	http.ListenAndServe(":"+w.WebServerPort, w.Router)
}
