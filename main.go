package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"kafka-gcp-connector/connector"
	"log"
	"os"
	"net/http"
)
var Port = fmt.Sprint(":"+ os.Getenv("PORT"))
func main()  {
	// Get Port From Enviornment variable
	if Port == "" {
		// If Port empty. than use default port:3000
		Port = ":3000"
	}
	// Init
	connector.InitFunc()
	// Status Check
	rh := chi.NewMux()
	rh.Get("/status", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("status:health"))
	})
	if err := http.ListenAndServe(Port, rh);err != nil {
		log.Printf("Start health check failed:%s", err.Error())
		return
	}
}