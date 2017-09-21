package main

import (
	"encoding/json"
	"net/http"
)

type Signup struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastname"`
}

func signup(w http.ResponseWriter, req *http.Request) {
	var info Signup
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	err := decoder.Decode(&info)

	if err != nil {
		http.Error(w, err.Error(), 400)

		return
		//		log.Fatal(err)
	}
	w.WriteHeader(200)
	w.Write([]byte("Successful"))
}
