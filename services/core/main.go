package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func show_time(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current Time: %s", time.Now())
}

func main() {
	var port = 6333
	http.HandleFunc("/", show_time)

	fmt.Println("Listening on http://localhost:" + strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
