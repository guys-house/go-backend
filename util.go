package main

import (
	"net/http"
	"encoding/json"
)

func CheckErr(err error){
	if err != nil {
		panic(err)
	}
}

func WriteSuccessOrFailure(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{
			Message: "error",
		})
		panic(err)
	} else {
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{
			Message: "success",
		})
	}
}