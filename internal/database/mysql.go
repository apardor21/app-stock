package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error

	dsn := "root:@tcp(localhost:3306)/stocksdb"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error abriendo la BD:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a la BD:", err)
	}

	log.Println("Conectado a MySQL correctamente")
}
