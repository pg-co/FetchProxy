package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/pg-co/FetchProxy/pkg/middleware"
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

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	resp, err := client.Get(payload.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.WriteHeader(resp.StatusCode)
	if _, err := io.Copy(w, resp.Body); err != nil {
		log.Printf("Failed to copy response body: %v", err)
		http.Error(w, "Failed to copy response body", http.StatusInternalServerError)
		return
	}
}


func main(){
	var (
		host = flag.String("host", "localhost", "Host of the database")
		port = flag.String("port", "8767", "Port of the database")
		auth = flag.String("auth", "user:password", "Basic auth credentials")
	)
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)

	mux := http.NewServeMux()
	mux.Handle("/proxy", middleware.Authenticate(*auth, http.HandlerFunc(executeFetch)))
	
	log.Printf("Server listening on: %v", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}