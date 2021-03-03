package main

import (
	"fmt"
	"goworkspace/apis/class_api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/apis/class/findall", class_api.FindAll).Methods("GET")
	router.HandleFunc("/apis/class/findemail", class_api.FindEmail).Methods("GET")

	err := http.ListenAndServe(":8001", router)
	if err != nil {
		fmt.Println(err)
	}
}
