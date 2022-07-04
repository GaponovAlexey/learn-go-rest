package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("hello%s\n", name)))
}

func main() {
	router := *httprouter.New()
	router.GET("/", indexHandler)

	listener, err := new.Listen("tcp", ":3000")
	if err != nil {panic(err)}
if err != nil {panic(err)}
}
