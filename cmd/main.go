package main

import (
	"net/http"
	"os"
	"log/slog"
)

func main() {

	port := ":8080"

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger = logger.With("app", "data-exchange-event","func", "main", "user", "user_id", "file_id", "file_id")

	logger.Info("Hello form the event app")

	logger.Debug("Starting server", "port", port)

	http.HandleFunc("/", index)

	_ = http.ListenAndServe(port, nil)

	// How to log stuff out:

	// 1. fmt 
	// 2. log
	// 3. using a library package
	// 4. slog

}//. 

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello Event"))
}// .index