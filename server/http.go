package server

import (
	"net/http"
	"os"
	"receipt-processor-challenge/api/rest"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	*mux.Router
	Address string
}

// New setups & returns a server
func New() *Server {
	router := mux.NewRouter()
	address := "localhost:8080"
	s := Server{Router: router, Address: address}
	router.HandleFunc("/receipts/process", rest.ProcessReceipts).
		Methods(http.MethodPost)
	router.HandleFunc("/receipts/{id}/points", rest.GetPointsForReceipt).
		Methods(http.MethodGet)
	return &s
}

func (s Server) ServeHTTP() {
	loggedRouter := handlers.LoggingHandler(os.Stdout, s.Router)
	srv := &http.Server{
		Handler: loggedRouter,
		Addr:    s.Address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Minute,
		ReadTimeout:  time.Minute,
	}
	logrus.Fatal(srv.ListenAndServe())
}
