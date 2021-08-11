package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func ServeHome(w http.ResponseWriter,r *http.Request)  {
	http.ServeFile(w,r,"home.html")
}


func main() {
	r:=mux.NewRouter()
	r.HandleFunc("/",ServeHome)
	http.ListenAndServe(":8081",r)
}