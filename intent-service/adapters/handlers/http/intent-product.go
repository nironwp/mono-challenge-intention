package handler

import (
	"encoding/json"
	"intent-service/domain/entities"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeIntentProductHandler(r *mux.Router, n *negroni.Negroni, service entities.IntentProductService) {
	r.Handle("/intent-product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")
}

func createProduct(service entities.IntentProductService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var productDto entities.NewIntentProductDto
		err := json.NewDecoder(r.Body).Decode(&productDto)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))

			return
		}

		productDto.ID = primitive.NewObjectID()

		product_make, err := entities.NewIntentProduct(productDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))

			return
		}
		product, err := service.Create(product_make)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}
