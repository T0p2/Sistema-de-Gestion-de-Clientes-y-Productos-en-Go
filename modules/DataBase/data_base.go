//go get -u github.com/mattn/go-sqlite3

package prueba

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Prueba_con_bd() *sql.DB {
	// Abrir la conexión a la base de datos (creará un archivo de base de datos SQLite en el directorio actual)
	db, err := sql.Open("sqlite3", "./clientes_productos.db")
	if err != nil {
		log.Fatal(err)
	}

	//tiramos un ping a la base de datos si existe.
	call := db.Ping()
	fmt.Print(call)
	//if call != nil {

	// Crear la tabla de clientes
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS clientes (
			idCliente INTEGER PRIMARY KEY AUTOINCREMENT,
			nombre VARCHAR(50),
			apellido VARCHAR(50),
			IdProducto INTEGER,
			FOREIGN KEY(IdProducto) REFERENCES productos(idProducto)
		)
	`)
	if err != nil {
		fmt.Print("ai")
		log.Fatal(err)
	}

	// Crear la tabla de productos
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS productos (
			idProducto INTEGER PRIMARY KEY AUTOINCREMENT,
			nombre VARCHAR(50),
			descripcion VARCHAR(50)
		)
	`)
	if err != nil {
		fmt.Print("eu")
		log.Fatal(err)
	}

	return db

}
