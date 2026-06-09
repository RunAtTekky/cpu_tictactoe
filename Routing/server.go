package routing

import (
	"fmt"
	"log"
	"net/http"
)

func Serve() {
	http.HandleFunc("/hello", Hello_handler)
	http.HandleFunc("/place", Place_handler)

	PORT := 8090
	log.Printf("Server listening on port %d\n", PORT)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		log.Fatalf("Server error %v", err)
	}
}
