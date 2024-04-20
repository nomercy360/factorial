package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /factorial/{number}", func(w http.ResponseWriter, r *http.Request) {
		number := r.PathValue("number")

		n, err := strconv.Atoi(number)

		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}

		result := factorial(n)

		_, _ = w.Write([]byte(strconv.Itoa(result)))
	})

	log.Println("Server running on port 8080")

	_ = http.ListenAndServe(":8080", router)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
