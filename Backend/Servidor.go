package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Valores struct {
	Valores []Valor `json:"Operaciones"`
}

type Valor struct {
	ID        primitive.ObjectID `bson:"_id"`
	Primero   float64            `bson:"Primero"`
	Segundo   float64            `bson:"Segundo"`
	Operacion string             `bson:"Operacion"`
	Resultado float64            `bson:"Resultado"`
	Fecha     time.Time          `bson:"Fecha"`
}

var collection *mongo.Collection
var ctx = context.TODO()

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ObtenerOperacion", ObtenerOperacion).Methods("POST")
	router.HandleFunc("/RegresarOperaciones", MandarOperaciones).Methods("GET")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":3000", handler))
}

func ObtenerOperacion(w http.ResponseWriter, r *http.Request) {
	var valores Valor

	//Traer informacion De Frontend Y Transformarla a struct
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type", "applicattion/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.Unmarshal(body, &valores)
	_ = json.NewEncoder(w).Encode(valores)

	//Calcular operaciones y obtener tiempo actual
	if valores.Operacion == "+" {
		valores.Resultado = valores.Primero + valores.Segundo
	}
	if valores.Operacion == "-" {
		valores.Resultado = valores.Primero - valores.Segundo
	}
	if valores.Operacion == "*" {
		valores.Resultado = valores.Primero * valores.Segundo
	}
	if valores.Operacion == "/" {
		if int(valores.Segundo) != 0 {
			valores.Resultado = valores.Primero / valores.Segundo
		} else {
			valores.Resultado = -0.000001
		}
	}

	//Convertir objeto a almacenar en coleccion
	operacion := bson.D{{"Primero", valores.Primero},
		{"Segundo", valores.Segundo},
		{"Operacion", valores.Operacion},
		{"Resultado", valores.Resultado},
		{"Fecha", time.Now()}}

	//Conectar con mongodb
	clientOptions := options.Client().ApplyURI("mongodb://192.168.0.7:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//Crear colleccion y base de datos si no existen y registrar en coleccion
	collection = client.Database("SO1_Practica1").Collection("Operaciones")
	_, err = collection.InsertOne(context.TODO(), operacion)
	if err != nil {
		panic(err)
	}
}

func MandarOperaciones(w http.ResponseWriter, _ *http.Request) {
	//Conectar con mongodb
	clientOptions := options.Client().ApplyURI("mongodb://192.168.0.7:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//Crear colleccion y base de datos si no existen y registrar en coleccion
	collection = client.Database("SO1_Practica1").Collection("Operaciones")

	//Traer todos los registros de coleccion
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	//Convierte todos los datos en bson
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	//Convierte datos bson en arreglo de operaciones para devolver a front como json
	var regreso Valores
	regreso.Valores = make([]Valor, len(results))
	w.Header().Set("Content-Type", "applicattion/json")
	w.WriteHeader(http.StatusAccepted)
	//Imprime todos los resultados
	for x, result := range results {
		var operacion Valor
		bsonBytes, _ := bson.Marshal(result)
		err = bson.Unmarshal(bsonBytes, &operacion)
		if err != nil {
			return
		}
		regreso.Valores[x] = operacion
	}

	//Codifica struct regreso en json
	_ = json.NewEncoder(w).Encode(regreso)
}
