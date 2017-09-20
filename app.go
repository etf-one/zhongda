package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type App struct {
	router     *mux.Router
	middleware http.Handler
}

func (a *App) initialize() {

	a.router = mux.NewRouter()
	a.initializeMiddlewares()
	a.initializeRoutes()

}

func (a *App) initializeRoutes() {

	a.router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome wefewf")
	})
}

func (a *App) initializeMiddlewares() {
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger()) // Includes some default middlewares
	n.UseHandler(a.router)
	a.middleware = n
}

func (a *App) Run(addr string) {

	http.ListenAndServe(addr, a.middleware)
}
