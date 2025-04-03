// made by recanman
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MoneroKon/wifireg/internal/config"
	"github.com/MoneroKon/wifireg/internal/registration"
)

func startServer() {
	registration.Init()
	http.HandleFunc("/register", registration.HandleAddUser)

	listenAddr := fmt.Sprintf(":%s", config.Port())
	fmt.Printf("Starting server on %s", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func main() {
	config.LoadEnv()
	startServer()
}
