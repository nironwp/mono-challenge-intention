package handler

import (
	"encoding/json"
	"intent-service/domain/entities"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeIntentHandler(r *mux.Router, n *negroni.Negroni, service entities.IntentService) {
	r.Handle("/intent", n.With(
		negroni.Wrap(createIntent(service)),
	)).Methods("POST", "OPTIONS")
	r.Handle("/intent", n.With(
		negroni.Wrap(getIntents(service)),
	)).Methods("GET", "OPTIONS")
}

func createIntent(service entities.IntentService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var intentDto entities.IntentDto

		err := json.NewDecoder(r.Body).Decode(&intentDto)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))

			return
		}

		for idx, _ := range intentDto.Itens {
			intentDto.Itens[idx].ID = primitive.NewObjectID()
		}

		intent_make, err := entities.NewIntent(intentDto)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))

			return
		}

		_, err = service.Create(intent_make)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonError(err.Error()))
			return
		}

	})
}

func getIntents(service entities.IntentService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		results, err := service.GetAll()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(results)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

	})
}
