package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"encoding/json"
	"math/big"
)

type Response struct {
	Result	string	`json:"result"`
}

// Successful API response
func okReq(w http.ResponseWriter, isJson bool, resp string) {
	if !isJson {
		fmt.Fprintf(w, "%s", resp)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `%s`, resp)
}

// Fail API response
func badReq(w http.ResponseWriter, isJson bool, err error) {
	if !isJson {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, `{"error":%q}`, err.Error())
}

// API Handel for trib()
func tribHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: tribHandler")
	vars := mux.Vars(r)
	n, err := strconv.Atoi(vars["n"])
	if err != nil {
		badReq(w, true, err)
		return
	}
	res,  err := json.Marshal(Response{Result: trib(big.NewInt(int64(n))).String()})
	if err != nil {
		badReq(w, true, err)
		return
	}
	okReq(w, true, string(res[:]))
}
