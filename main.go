package main

import (
	"github.com/go-chi/chi"
	"kafka-gcp-connector/connector"
	"log"
	"net/http"
)

func main()  {
	// Init
	connector.InitFunc()
	// Status Check
	rh := chi.NewMux()
	rh.Get("/status", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("status:health"))
	})
	if err := http.ListenAndServe(":3000", rh);err != nil {
		log.Printf("Start health check failed:%s", err.Error())
		return
	}
}