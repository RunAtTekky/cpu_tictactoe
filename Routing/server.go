package routing

import (
	"fmt"
	"net/http"
)

func Serve() {
	http.HandleFunc("/hello", Hello_handler)
	http.HandleFunc("/place", Place_handler)

	fmt.Println("Server listening on port 8090")

	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Printf("Server error %v", err)
	}
}
