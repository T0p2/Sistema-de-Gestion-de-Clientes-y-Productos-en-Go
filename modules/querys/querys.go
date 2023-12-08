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

func New_product_user(db *sql.DB, idProducto int, name string) {

	//inicio de transaccion, es un metodo que se asegura que todas las operaciones se cumplan o fallen todas.
	//lo implemento ya que sino la base de datos recibe muchas instrucciones SQL a la vez, y se desborda.

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer tx.Commit()

	// Verificar si el nombre del cliente ya existe en la tabla de clientes
	rowsName, err := tx.Query("SELECT nombre FROM clientes WHERE nombre = ?", name)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rowsName.Close()

	if !rowsName.Next() {
		fmt.Print("Error: El nombre del cliente no existe en la tabla de clientes.")
		return
	}

	// Verificar si el idProducto existe en la tabla de productos
	rows, err := tx.Query("SELECT idProducto FROM productos WHERE idProducto = ?", idProducto)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()

	if !rows.Next() {
		fmt.Println("Error: El idProducto no existe en la tabla de productos.")
		return
	}

	// Realizar la actualización en la tabla de clientes
	_, err = tx.Exec("UPDATE clientes SET idProducto = ? WHERE nombre = ?", idProducto, name)
	if err != nil {
		fmt.Println("Error al actualizar cliente:", err)
		return
	}

	fmt.Println("Producto asignado correctamente al cliente.")
}

func Load_product(db *sql.DB, name string, description string) {

	_, err := db.Exec("INSERT INTO productos (nombre, descripcion) VALUES (?, ?)", name, description)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print("Se ha cargado correctamente")
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
	defer rows.Close()

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

	rows, err := db.Query("SELECT idCliente, nombre, IdProducto FROM clientes")
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return list
	}
	defer rows.Close()

	for rows.Next() {
		var idCliente int
		var nombre sql.NullString
		var idProducto sql.NullInt16

		err := rows.Scan(&idCliente, &nombre, &idProducto)
		if err != nil {
			fmt.Println("Error al escanear la fila:", err)
			return list
		}

		//No soporta manejar nulos en caso de que el cliente no tenga productos, entonces los verificamos.
		var idProductoValid int
		if idProducto.Valid {
			idProductoValid = int(idProducto.Int16)
		} else {
			idProductoValid = -1
		}

		var nombreValue string
		if nombre.Valid {
			nombreValue = nombre.String
		} else {
			nombreValue = "NULL"
		}

		rowData := []string{fmt.Sprint(idCliente), nombreValue, fmt.Sprint(idProductoValid)}
		list = append(list, rowData)

	}

	return list
}

func CheckProduct(db *sql.DB, productName string) (bool, error) {
	// Realizar la consulta para verificar si el nombre del producto ya existe
	rows, err := db.Query("SELECT COUNT(*) FROM productos WHERE nombre = ?", productName)
	if err != nil {
		return false, nil
	}
	defer rows.Close()

	// Leer el resultado de la consulta
	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return false, nil
		}
	}

	// Si count es mayor que 0, el nombre del producto ya existe
	return count > 0, nil
}

func CheckClient(db *sql.DB, clientName string) (bool, error) {
	// Realizar la consulta para verificar si el nombre del producto ya existe
	rows, err := db.Query("SELECT COUNT(*) FROM clientes WHERE nombre = ?", clientName)
	if err != nil {
		return false, nil
	}
	defer rows.Close()

	// Leer el resultado de la consulta
	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return false, nil
		}
	}

	// Si count es mayor que 0, el nombre del producto ya existe
	return count > 0, nil
}
