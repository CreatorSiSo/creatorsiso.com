package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func get_time(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Current Time: %s %s", time.Now(), "\n")
	fmt.Fprintf(response, "test\n")
}

func get_name(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "A name")
}

func main() {
	var port = 6333
	http.HandleFunc("/", get_time)
	http.HandleFunc("/test", get_name)

	fmt.Println("Listening on http://localhost:" + strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
