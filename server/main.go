package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus" // library that helps with loging and monitoring
)

var PORT string = ":8080" // this is the port number for our endpoint
var CODE string = ""      // this will contain our generated code
var userCode string = ""  // this contains the code from client

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
	fmt.Printf("%s\n", "Def input was hit...empty!")
}

func openSafeEndPoint(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		log.Info("Someone hit the openSafe endpoint")
		// call our python program to generate password and send text to user

	default:
		fmt.Printf("%s\n", "Only accepting GET requests")
	}

}

func validateUserEndPoint(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		log.Info("Someone hit the validateEndPoint")
		req.ParseForm()
		userCode = req.FormValue("code")

		// if our code is not empty the we can compare it with the one provided by our generated password
		if len(CODE) > 0 {
			fmt.Printf("%s\n", userCode)
			// compare code with generated password and respond back to client on the auth status
		} else {
			fmt.Println("Erro with code!")
		}

	default:
		fmt.Printf("%s\n", "Only accepting POST requests")
	}
}
