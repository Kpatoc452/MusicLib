package main

import "net/http"

type Server struct{
	routes map[string]map[string]http.HandleFunc
}

func NewServer()

func (r *Server) addRoute(method, path string, handler http.HandlerFunc) {
    if r.routes[path] == nil {
        r.routes[path] = make(map[string]http.HandlerFunc)
    }
    r.routes[path][method] = handler
}

func main(){
	
}