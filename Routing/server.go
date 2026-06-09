package routing

import (
	"log"
	"net/http"
)

func Serve() {
	http.HandleFunc("/hello", Hello_handler)
	http.HandleFunc("/place", Place_handler)

	log.Println("Server listening on port 8090")

	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatalf("Server error %v", err)
	}
}
