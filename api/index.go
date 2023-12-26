package handler

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	fmt.Fprintf(w, "<a href=\"/api/json\">/api/json</a>")
}
