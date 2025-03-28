package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Configuración de la conexión a PostgreSQL
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "ecommerce306"
)

// Función para crear la base de datos si no existe
func createDatabase() {
	// Conectar a la base de datos principal
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		host, port, user, password)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error al conectar a PostgreSQL: %v", err)
	}
	defer db.Close()

	// Verificar si la base de datos existe
	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT FROM pg_database WHERE datname = $1)", dbname).Scan(&exists)
	if err != nil {
		log.Fatalf("Error al verificar la base de datos: %v", err)
	}

	// Crear la base de datos si no existe
	if !exists {
		_, err = db.Exec("CREATE DATABASE " + dbname)
		if err != nil {
			log.Fatalf("Error al crear la base de datos: %v", err)
		}
		fmt.Println("Base de datos 'ecommerce306' creada correctamente.")
	} else {
		fmt.Println("La base de datos 'ecommerce306' ya existe.")
	}
}

// Función para crear las tablas en la base de datos
func createTables() {
	// Conectar a la base de datos específica
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos 'ecommerce306': %v", err)
	}
	defer db.Close()

	// Definir las consultas para crear las tablas
	tables := map[string]string{
		"clientes": `
			CREATE TABLE IF NOT EXISTS clientes (
				id SERIAL PRIMARY KEY,
				nombre VARCHAR(255) NOT NULL,
				correo VARCHAR(255) UNIQUE NOT NULL,
				telefono VARCHAR(20),
				direccion TEXT
			)`,
		"productos": `
			CREATE TABLE IF NOT EXISTS productos (
				id SERIAL PRIMARY KEY,
				nombre VARCHAR(255) NOT NULL,
				descripcion TEXT,
				precio DECIMAL(10,2) NOT NULL CHECK (precio >= 0),
				stock INT NOT NULL CHECK (stock >= 0)
			)`,
		"ventas": `
			CREATE TABLE IF NOT EXISTS ventas (
				id SERIAL PRIMARY KEY,
				cliente_id INT NOT NULL,
				fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				total DECIMAL(10,2) NOT NULL CHECK (total >= 0),
				FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE CASCADE
			)`,
		"detalle_ventas": `
			CREATE TABLE IF NOT EXISTS detalle_ventas (
				id SERIAL PRIMARY KEY,
				venta_id INT NOT NULL REFERENCES ventas(id) ON DELETE CASCADE,
				producto_id INT NOT NULL REFERENCES productos(id),
				cantidad INT NOT NULL CHECK (cantidad > 0),
				subtotal DECIMAL(10,2) NOT NULL CHECK (subtotal >= 0)
			)`,
		"compras": `
			CREATE TABLE IF NOT EXISTS compras (
				id SERIAL PRIMARY KEY,
				fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				total DECIMAL(10,2) NOT NULL CHECK (total >= 0)
			)`,
		"detalle_compras": `
			CREATE TABLE IF NOT EXISTS detalle_compras (
				id SERIAL PRIMARY KEY,
				compra_id INT NOT NULL REFERENCES compras(id) ON DELETE CASCADE,
				producto_id INT NOT NULL REFERENCES productos(id),
				cantidad INT NOT NULL CHECK (cantidad > 0),
				costo DECIMAL(10,2) NOT NULL CHECK (costo >= 0)
			)`,
	}

	// Ejecutar la creación de cada tabla
	for name, query := range tables {
		_, err = db.Exec(query)
		if err != nil {
			log.Fatalf("Error al crear la tabla '%s': %v", name, err)
		}
		fmt.Printf("Tabla '%s' creada correctamente.\n", name)
	}
}

// Función principal
func main() {
	createDatabase() // Crear la base de datos
	createTables()   // Crear las tablas
}
