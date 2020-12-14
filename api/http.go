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

	mx.Handle("/api/signup", cors(signUpHandlerFunc)).Methods("POST", "OPTIONS")
	mx.HandleFunc("/api/login", loginHandler).Methods("POST")
	mx.Handle("/api/logout", cors(logoutHandlerFunc)).Methods("DELETE", "OPTIONS")
}
