package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Conectar a MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error al cerrar la conexión con MongoDB: %v", err)
		}
	}()

	// Verificar la conexión
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("No se pudo conectar a MongoDB: %v", err)
	}
	fmt.Println("Conexión exitosa a MongoDB.")

	// Seleccionar la base de datos
	db := client.Database("ecommerce3066478")

	// Crear colecciones
	clientes := db.Collection("clientes1")
	productos := db.Collection("productos1")

	// Insertar documentos de ejemplo
	clienteDoc := bson.D{
		{"nombre", "Juan Perez"},
		{"correo", "juanperez@gmail.com"},
		{"telefono", "555-1234"},
		{"direccion", "Calle Falsa 123"},
	}

	productoDoc := bson.D{
		{"nombre", "laptop"},
		{"descripcion", "laptop de alta gama"},
		{"precio", 1200},
		{"stock", 10},
	}

	_, err = clientes.InsertOne(context.TODO(), clienteDoc)
	if err != nil {
		log.Fatalf("Error al insertar cliente: %v", err)
	}

	_, err = productos.InsertOne(context.TODO(), productoDoc)
	if err != nil {
		log.Fatalf("Error al insertar producto: %v", err)
	}

	fmt.Println("Base de datos y colecciones creadas en MongoDB con documentos de ejemplo.")
}
