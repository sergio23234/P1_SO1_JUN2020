package main

import (
	"fmt"
	"net/http"

	info_cpu "./cpu_info"

	principal "./principal"

	ram_inf "./ram_inf"

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

	mux.HandleFunc("/ram", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var inf = ram_inf.ObtenerRAM()
		w.Write(inf)
	})

	mux.HandleFunc("/getprin", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var inf = principal.Principal()
		w.Write(inf)
	})
	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
