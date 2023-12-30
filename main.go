package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Message string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var requestBody RequestBody
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Некорректное JSON-сообщение", http.StatusBadRequest)
		return
	}

	if requestBody.Message == "" {
		http.Error(w, "Некорректное JSON-сообщение", http.StatusBadRequest)
		return
	}

	fmt.Printf("Сообщение от клиента: %s\n", requestBody.Message)

	response := Response{
		Status:  "success",
		Message: "Данные успешно приняты",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
