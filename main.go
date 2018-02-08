package main

import (
	"log"
	"net/http"
	"fmt"
)

func main(){
	router := NewRouter()
	fmt.Println("Server Start ----")
	log.Fatal(http.ListenAndServe(":9090", router))
}