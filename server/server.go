package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Starting Server muchtodo...")

	config := configSetup()

	// Health endpoint
	healthHandler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain")
		writer.WriteHeader(http.StatusOK)
		io.WriteString(writer, "OK")
	}

	// list get endpoint
	listHandler := func(writer http.ResponseWriter, request *http.Request) {
		list := getLists()
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(list)
	}

	// list update endpoint
	listUpdaterHandler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain")
		writer.WriteHeader(http.StatusCreated)
		io.WriteString(writer, addTaskToList("Added task from endpoint abc", "Demo Tasks"))
	}

	// register handlers
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/lists", listHandler)
	http.HandleFunc("/lists/update", listUpdaterHandler)

	// Start server
	server := &http.Server{
		Addr: config.Host + ":" + config.Port,
	}
	server.ListenAndServe()
}
