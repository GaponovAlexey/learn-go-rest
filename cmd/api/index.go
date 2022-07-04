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
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("name%v", name)))
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
