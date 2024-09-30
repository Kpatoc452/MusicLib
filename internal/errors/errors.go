package errors

import "net/http"

type Error struct{ 
}

func(e Error) notFound(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func(e Error) BadRequest(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400, bad request"))
}

func(e Error) MethodNotAllowed(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method not allowed"))
}