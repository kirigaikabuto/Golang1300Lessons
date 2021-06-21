package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type ErrorMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type MegaHelloResponse struct {
	Message string `json:"message"`
}

func MegaHelloHandler(param1, param2 string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//read params from path /hello/{config}
		data := mux.Vars(r)
		fmt.Println(data)
		config, ok := data[param1]
		if !ok {
			respondJSON(w, http.StatusBadRequest, &ErrorMessage{
				Message: "need parameter1",
				Status:  http.StatusBadRequest,
			})
			return
		}
		value, ok := data[param2]
		if !ok {
			respondJSON(w, http.StatusBadRequest, &ErrorMessage{
				Message: "need parameter2",
				Status:  http.StatusBadRequest,
			})
			return
		}
		message := config + value
		a := r.URL.Query().Get("a")
		if a != "" {
			message += a
		}
		respondJSON(w, 200, &MegaHelloResponse{Message: message})
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/hello/{config}/{value}").HandlerFunc(MegaHelloHandler("config", "value"))
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
