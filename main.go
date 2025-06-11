package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type Response struct {
    Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format(time.RFC3339)
    response := Response{
        Time: currentTime,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/time", timeHandler)

    fmt.Println("Server started at http://localhost:8795")
    if err := http.ListenAndServe(":8795", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}