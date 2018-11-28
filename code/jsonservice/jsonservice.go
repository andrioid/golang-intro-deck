package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type greetInput struct {
	Name string `json:"name"`
}

type greetOutput struct {
	Text string `json:"text"`
}

func HandleGreet(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input greetInput
	var output greetOutput
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	} else {
		output.Text = fmt.Sprintf("Hello %s", input.Name)
		json.NewEncoder(w).Encode(output)
	}
}

func main() {
	http.HandleFunc("/greet", HandleGreet)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to PARKPARK!")
	})

	http.ListenAndServe(":3022", nil)
}
