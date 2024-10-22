package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
)

type Payload struct {
	Url string `json:"url"`
}


func executeFetch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := http.Get(payload.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)	
}


func main(){
	var (
		host = flag.String("host", "localhost", "Host of the database")
		port = flag.String("port", "8767", "Port of the database")
	)
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)

	mux := http.NewServeMux()
	mux.HandleFunc("/proxy", executeFetch)
	
	log.Printf("Server listening on: %v", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}