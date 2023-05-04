package handler

import (
	"net/http"

	"github.com/augustopedro/hexagonal-go/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProduckHandler(router *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(request)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
	})
}
