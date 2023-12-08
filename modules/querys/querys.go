package query

import (
	"database/sql"
	"fmt"
	"log"
)

func Load_user(db *sql.DB, name string, last_name string) {

	_, err := db.Exec("INSERT INTO clientes (nombre, apellido)  VALUES (?, ?)", name, last_name)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print("Cargado correctamente")
	}
}

func Delete_user(db *sql.DB, name string) {
	_, err := db.Exec("DELETE FROM clientes WHERE nombre = ?", name)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Usuario eliminado correctamente")
	}
}

func Update_user(db *sql.DB, oldName string, newName string, newLastName string) {
	_, err := db.Exec("UPDATE clientes SET nombre = ?, apellido = ? WHERE nombre = ?", newName, newLastName, oldName)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Usuario actualizado correctamente")
	}
}

func New_product_user(db *sql.DB, id_producto int, name string) {

	// Verificar si el nombre del cliente ya existe en la tabla de clientes
	rowsName, err := db.Query("SELECT nombre FROM clientes WHERE nombre = ?", name)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rowsName.Close()

	var nombreFromDB string
	if rowsName.Next() {
		if err := rowsName.Scan(&nombreFromDB); err != nil {
			log.Fatal(err)
			return
		}
	} else {
		fmt.Println("Error: El nombre del cliente no existe en la tabla de clientes.")
		return
	}

	// Aca verificamos el tema del producto
	rows, err := db.Query("SELECT idProducto FROM productos WHERE idProducto = ?", id_producto)
	if err != nil {
		fmt.Print("aquui")
		log.Fatal(err)
		return
	}
	defer rows.Close()

	var idProductoFromDB int
	if rows.Next() {
		if err := rows.Scan(&idProductoFromDB); err != nil {
			log.Fatal(err)
			return
		}
	} else {
		fmt.Println("Error: El id_producto no existe en la tabla de productos.")
		return
	}

	//aca hacemos la actualizacion
	_, err1 := db.Exec("UPDATE clientes SET IdProducto = ? WHERE nombre = ?", id_producto, name)
	if err1 != nil {
		fmt.Print("No existe el nombre o id")
		log.Fatal(err1)
	} else {
		fmt.Print("Insertado correctamente.")
	}
}
func Load_product(db *sql.DB, name string, description string) {

	_, err := db.Exec("INSERT INTO productos (nombre, descripcion) VALUES (?, ?)", name, description)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print("Se ah cargado correctamente")
	}
}

func Delete_product(db *sql.DB, name string) {
	_, err := db.Exec("DELETE FROM productos WHERE nombre = ?", name)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Producto eliminado correctamente")
	}
}

func Update_product(db *sql.DB, oldName string, newName string, description string) {
	_, err := db.Exec("UPDATE productos SET nombre = ?, descripcion = ? WHERE nombre = ?", newName, description, oldName)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Producto actualizado correctamente")
	}
}

func Return_id_product(db *sql.DB) [][]string {
	var list [][]string
	//var idProducto int

	rows, err := db.Query("SELECT idProducto, nombre FROM productos")

	if err != nil {
		fmt.Print("NO HAY PRODUCTOS ERROR")
	}

	for rows.Next() {
		var idProducto int
		var nombre string

		err := rows.Scan(&idProducto, &nombre)
		if err != nil {
			// Manejar el error de escaneo
			fmt.Println("Error al escanear la fila:", err)
		}

		// Crear una nueva lista para cada fila y agregarla a la lista principal
		rowData := []string{fmt.Sprint(idProducto), nombre}
		list = append(list, rowData)
	}

	return list
}

func Return_id_client(db *sql.DB) [][]string {
	var list [][]string

	rows, err := db.Query("SELECT idCliente, nombre FROM clientes")
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return list
	}
	defer rows.Close()

	for rows.Next() {
		var idCliente int
		var nombre sql.NullString

		err := rows.Scan(&idCliente, &nombre)
		if err != nil {
			fmt.Println("Error al escanear la fila:", err)
			continue
		}

		var nombreValue string
		if nombre.Valid {
			nombreValue = nombre.String
		} else {
			nombreValue = "N/A" // o cualquier valor predeterminado que desees para NULL
		}

		rowData := []string{fmt.Sprint(idCliente), nombreValue}
		list = append(list, rowData)
	}

	return list
}
