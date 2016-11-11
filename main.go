package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const serverPORT = ":8000"

// Reply is a data structure for the server replies
type Reply struct {
	Hostname string
	ClientIP string
	Time     time.Time
}

func serverRESPONSE(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname() // Get the local hostname

	// Get the client's IP address.
	// RemoteAddr returns the client IP address with the port after a colon
	// We split the client IP + port based on colon(s) and only remove
	// after the last one, so we don't break IPv6
	clientsocket := r.RemoteAddr
	clientipMap := strings.Split(clientsocket, ":")
	clientipMap = clientipMap[:len(clientipMap)-1]
	clientip := strings.Join(clientipMap, ":")

	response := Reply{hostname, clientip, time.Now()} // Construct the response with the gathered data

	// Convert to json
	jsonRESPONSE, err := json.Marshal(response)
	if err != nil {
		log.Output(0, "json conversion failed")
	}

	io.WriteString(w, string(jsonRESPONSE)) // Send response back to client
}

func main() {
	http.HandleFunc("/", serverRESPONSE)
	http.ListenAndServe(serverPORT, nil)
}
