package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "{id: 1, title: 2}")
}

func Hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", params.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.GET("/hello/:name", Hello)

	listener, err := net.Listen("tcp", "learn-go-rest.vercel.app")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.Serve(listener))
}
