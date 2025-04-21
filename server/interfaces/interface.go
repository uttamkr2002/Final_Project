package interfaces

import (
	"net/http"
	"github.com/gorilla/mux"
)

// RouterInterface defines routing methods for abstraction.
type RouterInterface interface {
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request) //
}

// GorillaMuxRouter is an implementation of RouterInterface using Gorilla Mux.
type GorillaMuxRouter struct {
	muxRouter *mux.Router
}

// NewGorillaMuxRouter creates a new instance of GorillaMuxRouter.
func NewGorillaMuxRouter() *GorillaMuxRouter {
	return &GorillaMuxRouter{muxRouter: mux.NewRouter()}
}

type Nethttp struct {
	muxRouter *http.ServeMux
}

func NewNethttp() *Nethttp {
	return &Nethttp{muxRouter: http.NewServeMux()}
}

func (g *GorillaMuxRouter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) {
	g.muxRouter.HandleFunc(path, f).Methods(http.MethodPost)
}

func (g *GorillaMuxRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.muxRouter.ServeHTTP(w, r)
}

func (n *Nethttp) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) {
	n.muxRouter.HandleFunc(path, f)
}

func (n *Nethttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	n.muxRouter.ServeHTTP(w, r)
}
