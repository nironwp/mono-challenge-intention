package httpport

import (
	"fmt"
	handler "intent-service/adapters/handlers/http"
	"intent-service/domain/entities"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service entities.IntentService
}

func MakeNewWebServer(service entities.IntentService) *WebServer {
	return &WebServer{
		Service: service,
	}
}

func (w WebServer) Server() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeIntentHandler(r, n, w.Service)

	http.Handle("/", r)
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	fmt.Println("Server running on port "+server.Addr, " PID:", os.Getpid())
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
