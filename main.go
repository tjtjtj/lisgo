package main

import (
	"github.com/tjtjtj/lisgo/handler"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

const port = "8080"

func main() {
	router := mux.NewRouter()
	router.Methods("GET").HandlerFunc(handler.Get)
	router.Handle("/", router)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatal("err: %v", err)
	}
}
