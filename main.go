package main

import (
    "encoding/json"
    "net/http"
    "strconv"
)

// Response struct untuk format JSON
type Response struct {
    Message string `json:"message"`
}

// Handler untuk endpoint root
func helloHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{Message: "Hello, World!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// Handler untuk endpoint penjumlahan
func addHandler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    a := query.Get("a")
    b := query.Get("b")

    num1, err1 := strconv.Atoi(a)
    num2, err2 := strconv.Atoi(b)

    if err1 != nil || err2 != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    sum := num1 + num2
    response := Response{Message: "Sum: " + strconv.Itoa(sum)}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/add", addHandler)

    // Menjalankan server di port 9090
    http.ListenAndServe(":9090", nil)
}
