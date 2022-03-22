package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tan4ek/openvpn-http-api/ovpn"
)

func Run(port string) {
	initRoutes()

	log.Printf("Starting server on port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

// GET  /ovpn-config?clientId=<user>
// POST /ovpn-config?clientId=<user>&password=<password>
func initRoutes() {
	log.Print("Initializing routes")

	http.HandleFunc("/ovpn-config", ovpnConfigRouteHandler)
}

func ovpnConfigRouteHandler(w http.ResponseWriter, req *http.Request) {
	clientId, err := getQueryParam("clientId", req)

	if err != nil {
		badRequest(w, err)
		return
	}

	if !validateClientId(clientId) {
		badRequest(w, fmt.Errorf("clientId must be at least 4 characters long"))
		return
	}

	if req.Method == "POST" {
		log.Print("POST /ovpn-config")

		password, err := getQueryParam("password", req)

		if err != nil {
			badRequest(w, err)
			return
		}

		if !validatePassword(password) {
			badRequest(w, fmt.Errorf("password must be at least 4 characters long"))
			return
		}

		err = ovpn.GenerateClientCerts(clientId, password)

		if err != nil {
			conflict(w, fmt.Errorf("client certs generation failed: %s", err))
			return
		}
	}

	config, err := ovpn.GenerateClientConfig(clientId)

	if err == nil {
		fmt.Fprint(w, config)
	} else {
		notFound(w, err)
	}
}

func getQueryParam(name string, req *http.Request) (string, error) {
	value, ok := req.URL.Query()[name]

	if ok {
		return value[0], nil
	} else {
		return "", fmt.Errorf("no query param \"%s\" found", name)
	}
}

func conflict(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusConflict)
	fmt.Fprint(w, err)
}

func badRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, err)
}

func notFound(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, err)
}

func validateClientId(value string) bool {
	return len(value) >= 4
}

func validatePassword(value string) bool {
	return len(value) >= 4
}
