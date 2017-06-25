package service

import (
	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
	"github.com/gorilla/mux"
	"net/http"
)

func NewServer() *negroni.Negroni {
	formater := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formater)
	n.UseHandler(mx)

	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/test", testHandler(formatter)).Methods(http.MethodGet)
}

func testHandler(formater *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formater.JSON(w, http.StatusOK, struct{Test string} {"This is a test"})
	}
}
