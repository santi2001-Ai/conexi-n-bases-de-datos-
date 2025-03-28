package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Importar el driver MySQL
)

func main() {
	// Configurar la conexión a MySQL
	dsn := "root:@tcp(localhost:3306)/"

	// Crear la conexión
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Error al conectar a MySQL:", err)
	}
	defer db.Close()

	// Verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatal("❌ No se pudo conectar a MySQL:", err)
	}

	fmt.Println("✅ Conexión exitosa a MySQL")
}
