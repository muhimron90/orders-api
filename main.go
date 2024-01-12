package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}
	// listen server
	// if error was found or not null (nil)
	// print errror
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to liesten to server", err)
	}
}
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))

}
