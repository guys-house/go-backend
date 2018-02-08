package main

import (
	"net/http"
	"io/ioutil"
	"io"
	"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

func CreateRecordIndex(w http.ResponseWriter, r *http.Request) {
	var record Record
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &record); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	err = CreateRecord(record)
	WriteSuccessOrFailure(w, err)
}

func RemoveRecordIndex(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, _ := strconv.Atoi(id)
	fmt.Println(idInt)
	err := RemoveRecord(idInt)
	WriteSuccessOrFailure(w, err)
}