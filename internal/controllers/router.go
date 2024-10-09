package controllers

import (
	"fmt"
	"net/http"
)

type Router struct{
	handler *Handler
}

func NewRouter(handler *Handler) *Router{
	return &Router{
		handler: handler,
	}
}

func(rout *Router) Hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello")
}

func(rout *Router) Song(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case http.MethodGet:
			// fmt.Fprintf(w, "hello world!")
			rout.handler.GetSong(w, r)
		case http.MethodPost:
			rout.handler.CreateSong(w, r)
		case http.MethodPut:
			rout.handler.UpdateSong(w, r)
		case http.MethodDelete:
			rout.handler.DeleteSong(w, r)
	}
}