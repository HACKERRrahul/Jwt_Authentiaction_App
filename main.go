package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
func InitializeRouter(){
	r:=mux.NewRouter()
	r.HandleFunc("/login",Login).Methods("POST")
	r.HandleFunc("/user/get",GetUser).Methods("GET")
	r.HandleFunc("/user/create",CreateUser).Methods("POST")
	r.HandleFunc("/refresh",Refresh).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090",r))
	
}
func main()  {
	
	InitialMigration()
	InitializeRouter()

}
