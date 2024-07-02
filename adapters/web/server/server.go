package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Math2121/hexago/adapters/web/handler"
	"github.com/Math2121/hexago/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	router := mux.NewRouter()
	negao := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(router, negao, w.Service)
	http.Handle("/", router)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
