package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("SiteSathi API starting on :8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("SiteSathi API running"))
	})
	http.ListenAndServe(":8080", nil)
}
