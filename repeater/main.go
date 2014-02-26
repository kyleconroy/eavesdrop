package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

    ping := func() {
        http.Get("http://localhost:8080/bar")
    }

	go func() {
		for {
			log.Println("Ping")
            ping()
			time.Sleep(time.Second * 1)
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
