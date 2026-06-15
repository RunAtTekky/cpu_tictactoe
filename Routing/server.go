package routing

import (
	"fmt"
	"log"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("content-type", "application/json")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Serve() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", Hello_handler)
	mux.HandleFunc("/validate", Validate_handler)
	mux.HandleFunc("/place", Place_handler)
	mux.HandleFunc("/has_won", Has_won_handler)
	mux.HandleFunc("/game_over", Game_over_handler)
	mux.HandleFunc("/restart", Restart_handler)

	handler := corsMiddleware(mux)

	PORT := 8090
	log.Printf("Server listening on port %d\n", PORT)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), handler); err != nil {
		log.Fatalf("Server error %v", err)
	}
}
