package simpleserver

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type APIServer struct {
// 	port string
// }

// func NewServer() *APIServer {
// 	return &APIServer{
// 		port: ":8233",
// 	}
// }

func start() {
	mux := mux.NewRouter()

	mux.HandleFunc("/checkmarx", handler)

	http.ListenAndServe(":8233", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		io.WriteString(w, "oops")
		return
	}

	w.Header().Set("Response-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data, err := io.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, "oops")
		return
	}

	io.WriteString(w, string(data))

	err = json.NewEncoder(w).Encode(map[string]any{"age": 10})
	if err != nil {
		log.Fatal(err)
	}
	return
}
