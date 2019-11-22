package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/quotes", quotes)

	fmt.Printf("Server listening on port 8080\n")
	log.Panic(http.ListenAndServe(":8080", nil))

	
}

func homePage(writer http.ResponseWriter, reqest *http.Request) {
	_, err := fmt.Fprintf(writer, "Welcome to the inspirational quote API homepage\n")
	if err != nil {
		log.Panic(err)
	}
}

func quotes(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		newQuote(writer, request)
		return
	}

	if request.Method == http.MethodGet {
		getRandomQuote(writer)
		return
	}

	writeResponseOrPanic(writer, "Invalid request method\n")
}

// insert a new quote into database
func newQuote(writer http.ResponseWriter, request *http.Request) {
	quote, err := NewQuoteFromRequest(request)
	if err != nil {
		writeResponseOrPanic(writer, fmt.Sprintf("Error: unable to create a quote from request data.\nmessage: %s\n", err.Error()))
		return
	}
	err = quote.storeInDatabase()
	if err != nil {
		writeResponseOrPanic(writer, fmt.Sprintf("error while storing quote in database.\nmessage: %s\n", err.Error()))
		return
	}

	writeResponseOrPanic(writer, fmt.Sprintf("Quote added: \"%s\"\n", quote.Quote))
}

// Get a random quote from database
func getRandomQuote(writer http.ResponseWriter) {
	quoteStruct, err := RandomQuoteFromDatabase()
	if err != nil {
		writeResponseOrPanic(writer, fmt.Sprintf("Error: unable to get quote from database\n%s\n", err.Error()))
		return
	}
	writeResponseOrPanic(writer, fmt.Sprintf("{\"quote\": \"%s\"}\n", quoteStruct.Quote))
}

// Will write a response using the http.ResponseWriter. If it fails it will panic.
func writeResponseOrPanic(writer http.ResponseWriter, message string) {
	_, err := fmt.Fprint(writer, message)
	if err != nil {
		log.Panic(err)
	}
}

func writeJson(writer http.ResponseWriter, data interface{}, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Panic(err)
	}

	writeResponseOrPanic(writer, string(jsonBytes))
}
