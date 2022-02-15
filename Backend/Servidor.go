package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Valores struct {
	Primero   float64 `json:"Primero"`
	Segundo   float64 `json:"Segundo"`
	Operacion string  `json:"Operacion"`
}

var collection *mongo.Collection
var ctx = context.TODO()

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ObtenerOperacion", ObtenerOperacion).Methods("POST")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":3000", handler))
}

func ObtenerOperacion(w http.ResponseWriter, r *http.Request) {
	var valores Valores
	var resultado float64

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type", "applicattion/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.Unmarshal(body, &valores)
	_ = json.NewEncoder(w).Encode(valores)

	if valores.Operacion == "+" {
		resultado = valores.Primero + valores.Segundo
	}
	if valores.Operacion == "-" {
		resultado = valores.Primero - valores.Segundo
	}
	if valores.Operacion == "*" {
		resultado = valores.Primero * valores.Segundo
	}
	if valores.Operacion == "/" {
		if int(valores.Segundo) != 0 {
			resultado = valores.Primero / valores.Segundo
		} else {
			resultado = -0.000001
		}
	}
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("SO1_Practica1").Collection("Operaciones")

	println(valores.Primero)
	println(valores.Operacion)
	println(valores.Segundo)
	println(resultado)
	println(fecha)
	println()
}
