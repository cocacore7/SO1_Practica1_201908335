package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ObtenerSuma", ObtenerSuma).Methods("GET")
	router.HandleFunc("/EnviarSuma", EnviarSuma).Methods("POST")

	router.HandleFunc("/ObtenerResta", ObtenerResta).Methods("GET")
	router.HandleFunc("/EnviarResta", EnviarResta).Methods("POST")

	router.HandleFunc("/ObtenerMultiplicacion", ObtenerMultiplicacion).Methods("GET")
	router.HandleFunc("/EnviarMultiplicacion", EnviarMultiplicacion).Methods("POST")

	router.HandleFunc("/ObtenerDivision", ObtenerDivision).Methods("GET")
	router.HandleFunc("/EnviarDivision", EnviarDivision).Methods("POST")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":3000", handler))
}

func ObtenerSuma(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "applicattion/json")
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode("Datos Suma Recibidos")
}

func EnviarResta(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "applicattion/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode("Enviar Suma")
}

func ObtenerResta(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "applicattion/json")
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode("Datos Resta Recibidos")
}

func EnviarMultiplicacion(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "applicattion/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode("Enviar Suma")
}

func ObtenerMultiplicacion(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "applicattion/json")
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode("Datos Multiplicacion Recibidos")
}

func EnviarDivision(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "applicattion/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode("Enviar Suma")
}

func ObtenerDivision(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "applicattion/json")
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode("Datos Division Recibidos")
}

func EnviarSuma(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "applicattion/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode("Enviar Suma")
}
