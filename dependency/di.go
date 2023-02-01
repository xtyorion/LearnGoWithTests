package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func HttpGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func CliGreeterHandler() {
	Greet(os.Stdout, "Elodie")
}

func main() {
	CliGreeterHandler()
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(HttpGreeterHandler)))
}
