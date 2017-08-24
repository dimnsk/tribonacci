package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// Run HTTP server
// Listen on 8081 port and endpoint "/v1/trib/{n}" where n - parameter
func startServer() {
	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.HandleFunc("/v1/trib/{n}", tribHandler).Name("/v1/trib/{n}").Methods("GET")
	http.Handle("/", mainRouter)

	fmt.Println("Starting service on 8081 port...")
	err := http.ListenAndServe(":8081", mainRouter)
	if err != nil {
		fmt.Println("Could not start server on 8081 port. " + err.Error())
	}
}

func main() {
	startServer()
}
