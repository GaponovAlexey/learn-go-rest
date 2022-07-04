package handlers

import "github.com/julienschmidt/httprouter"

//method register
type Handler interface {
	Register(router *httprouter.Router)
}
