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
	clientId := getQueryParam("clientId", req)

	switch req.Method {
	case "GET":
		log.Print("GET /ovpn-config")
	case "POST":
		log.Print("POST /ovpn-config")

		password := getQueryParam("password", req)
		ovpn.GenerateClientCerts(clientId, password)
	}

	config, err := ovpn.GenerateClientConfig(clientId)

	if err == nil {
		fmt.Fprint(w, config)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err)
	}
}

func getQueryParam(name string, req *http.Request) string {
	value, ok := req.URL.Query()[name]

	if ok {
		return value[0]
	} else {
		log.Fatal("No query param found")

		return ""
	}
}
