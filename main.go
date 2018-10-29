package main

import (
	"tjtjtj/lisgo/handler"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

const port = "8080"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/articles", handler.Articles)
	router.Methods("GET").HandlerFunc(handler.Get)

	//router.PathPrefix("/").Methods("GET").Subrouter().HandleFunc("/", handler.Get)
	//router.Methods("GET").Subrouter().HandleFunc("/", handler.Get)
	router.HandleFunc("/", handler.Top)

	router.Handle("/", router)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatal("err: %v", err)
	}
}
