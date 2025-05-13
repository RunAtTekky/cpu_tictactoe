package main

import (
	"fmt"
	"net/http"
)

func hello_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func main() {
	http.HandleFunc("/hello", hello_handler)
	fmt.Println("Server listening on port 8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Printf("Server error %v", err)
	}

}
