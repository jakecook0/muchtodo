package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ListUpdate struct {
	Listname string
	Task     string
}

type NewList struct {
	Listname string
}

func main() {
	fmt.Println("Starting Server muchtodo...")

	config := configSetup()

	// Health handler
	healthHandler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain")
		writer.WriteHeader(http.StatusOK)
		io.WriteString(writer, "OK")
	}

	// get lists handler
	listHandler := func(writer http.ResponseWriter, request *http.Request) {
		list := getLists()
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(list)
	}

	// update lists handler
	listUpdatHandler := func(writer http.ResponseWriter, request *http.Request) {
		// listname := strings.Split(request.RequestURI, "?")
		var newtask ListUpdate
		err := json.NewDecoder(request.Body).Decode(&newtask)
		if err != nil {
			println("Error decoding JSON: %s\n", err)
		}

		err = addTaskToList(newtask.Listname, newtask.Task)

		if err != nil {
			writer.Header().Set("Content-Type", "text/plain")
			writer.WriteHeader(http.StatusBadRequest)
			io.WriteString(writer, "Error - bad request body")
		}

		writer.Header().Set("Content-Type", "text/plain")
		writer.WriteHeader(http.StatusCreated)
		io.WriteString(writer, "Success - List updated with new task")
	}

	//create a new list
	listCreateHandler := func(writer http.ResponseWriter, request *http.Request) {
		listname := request.URL.Query().Get("listname")
		err := createList(listname)

		if err != nil {
			writer.Header().Set("Content-Type", "text/plain")
			writer.WriteHeader(http.StatusBadRequest)
			io.WriteString(writer, "Error - bad request body")
		}

		writer.Header().Set("Content-Type", "text/plain")
		writer.WriteHeader(http.StatusCreated)
		io.WriteString(writer, "Success - List updated with new task")
	}

	// register handlers
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/lists", listHandler)
	http.HandleFunc("/lists/update", listUpdatHandler)
	http.HandleFunc("/lists/create", listCreateHandler)

	// Start server
	server := &http.Server{
		Addr: config.Host + ":" + config.Port,
	}
	server.ListenAndServe()
}
