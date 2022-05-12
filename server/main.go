package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus" // library that helps with loging and monitoring
)

var PORT string = ":8080"

func main() {
	fmt.Printf("%s\n", "hello world!")
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting Server...")
	http.HandleFunc("/", defEndPoint)
	http.HandleFunc("/api/v2/open_safe", openSafeEndPoint)
	http.HandleFunc("/api/v2/validate", validateUserEndPoint)
	http.ListenAndServe(PORT, nil)
}

func defEndPoint(w http.ResponseWriter, req *http.Request) {

}

func openSafeEndPoint(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		log.Info("Someone hit the openSafe endpoint")
	default:
		fmt.Printf("%s\n", "Only accepting GET requests")
	}

}

func validateUserEndPoint(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		log.Info("Someone hit the validateEndPoint")
	default:
		fmt.Printf("%s\n", "Only accepting GET requests")
	}
}
