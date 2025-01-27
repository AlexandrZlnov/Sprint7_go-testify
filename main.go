package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var cafeList = map[string][]string{
	"moscow": {"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ", ")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

func main() {
	address := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/cafe/", mainHandle)

	srv := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	fmt.Printf("Starting server on %s\n", address)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("failed to listen and serve: %s\n", err)
	}

}
