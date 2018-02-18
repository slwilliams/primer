package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func startServer() *http.Server {

	r := mux.NewRouter()
	r.HandleFunc("/", getPrimeEndpoint)
	r.HandleFunc("/{max:[0-9]+}", getPrimeArgEndpoint)
	r.HandleFunc("/_ah/health", healthCheckHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", defaultPort),
		Handler:      handlers.CombinedLoggingHandler(os.Stdout, r),
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	return srv

}

func getPrimeEndpoint(w http.ResponseWriter, req *http.Request) {
	writeResponse(defaultMaxNumber, w)
}

func getPrimeArgEndpoint(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	maxNum, err := strconv.Atoi(vars["max"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		writeResponse(maxNum, w)
	}
}

func writeResponse(maxNum int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ServiceResponse{
		ID:    getUUIDv4(),
		Host:  hostname,
		Ts:    time.Now().UTC().String(),
		Prime: *getPrime(maxNum),
	})

}

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}
