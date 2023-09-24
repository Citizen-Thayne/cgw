package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

var counter = 0

func incr(w http.ResponseWriter, r *http.Request) {
	counter = counter + 1
	io.WriteString(w, strconv.Itoa(counter))
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, strconv.Itoa(counter))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/api/counter/incr", incr)
	http.HandleFunc("/api/counter", getCounter)
	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
