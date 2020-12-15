package api

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

//NewServer configure and returns a Server
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	signUpHandlerFunc := http.HandlerFunc(signUpHandler)
	logoutHandlerFunc := http.HandlerFunc(logoutHandler)
	loginHandlerFunc := http.HandlerFunc(loginHandler)

	mx.Handle("/api/signup", cors(signUpHandlerFunc)).Methods("POST", "OPTIONS")
	mx.Handle("/api/login", cors(loginHandlerFunc)).Methods("POST", "OPTIONS")
	mx.Handle("/api/logout", cors(logoutHandlerFunc)).Methods("DELETE", "OPTIONS")
}
