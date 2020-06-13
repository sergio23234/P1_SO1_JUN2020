package main

import (
	"fmt"
	"net/http"

	info_cpu "./cpu_info"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hola\": \"mundo\"}"))
	})

	mux.HandleFunc("/cpu", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		inf := fmt.Sprintf("%f", info_cpu.GetPorcentajeUso())
		w.Write([]byte("{\"cpu_info\": \"" + inf + "\"}"))
	})

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
