package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
func InitializeRouter(){
	r:=mux.NewRouter()
	//r.HandleFunc("/users",GetUsers).Methods("GET")
	r.HandleFunc("/login",Login).Methods("POST")
	r.HandleFunc("/user",GetUser).Methods("GET")
	r.HandleFunc("/users",CreateUser).Methods("POST")
	//r.HandleFunc("/users/{id}",UpdateUser).Methods("PUT")
	//r.HandleFunc("/users/{id}",DeleteUser).Methods("DELETE")
	r.HandleFunc("/refresh",Refresh).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090",r))
	
}
func main()  {
	
	InitialMigration()
	InitializeRouter()

}