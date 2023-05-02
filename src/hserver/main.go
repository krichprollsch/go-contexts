package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Simulate a long operation, like an HTTP request.
		select {
		case <-ctx.Done():
			fmt.Println("cancel", r.RemoteAddr)
			return
		case <-time.After(5 * time.Second):
			// nothing to do here.
		}
		fmt.Fprintln(w, "done")
	})

	fmt.Fprintln(os.Stderr, http.ListenAndServe("127.0.0.1:4567", nil))
}
