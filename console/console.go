package console

import (
	"database/sql"
	"fmt"
	"log"
	functions "main/modules/querys"
)

type Cliente struct {
	idCliente int
	Nombre    string
	Apellido  string
}

type Producto struct {
	idProducto  int
	Nombre      string
	Descripcion string
}

func Talk(db *sql.DB) {
	operation := getUserInput("¿Qué operación desea realizar? (load_user, load_product, delete_user, delete_product, update_user, update_products, list_user, list_product): ")

	switch operation {
	case "load_user":
		name := getUserInput("Ingrese nombre: ")
		last_name := getUserInput("Ingrese apellido: ")
		functions.Load_user(db, name, last_name)
	case "load_product":
		name := getUserInput("Ingrese nombre del producto: ")
		description := getUserInput("Ingrese la descripcion del producto : ")
		functions.Load_product(db, name, description)
	case "delete_user":
		name := getUserInput("Ingrese nombre del usuario a eliminar: ")
		functions.Delete_user(db, name)
	case "delete_product":
		name := getUserInput("Ingrese nombre del producto a eliminar: ")
		functions.Delete_product(db, name)
	case "update_user":
		name_old := getUserInput("Ingrese el nombre del usuario que quiere actualizar: ")
		name_new := getUserInput("Ingrese el nuevo nombre: ")
		new_last_name := getUserInput("Ingrese el nuevo apellido: ")
		functions.Update_user(db, name_old, name_new, new_last_name)
	case "update_product":
		name_old := getUserInput("Ingrese el nombre del producto que quiere actualizar: ")
		name_new := getUserInput("Ingrese el nuevo nombre: ")
		description := getUserInput("Ingrese la nueva descripcion:  ")
		functions.Update_product(db, name_old, name_new, description)
	case "list_user":
		list_users(db)
	case "list_product":
		list_product(db)
	default:
		fmt.Println("Operación no válida")
	}
}

func getUserInput(message string) string {
	var input string
	fmt.Print(message)

	fmt.Scan(&input)
	return input

}

func list_users(db *sql.DB) {

	rows, err := db.Query("SELECT * FROM clientes")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print("Lista de Usuarios: ")
		for rows.Next() {
			var c Cliente
			err := rows.Scan(&c.idCliente, &c.Nombre, &c.Apellido)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(c)
		}

	}
}

func list_product(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM productos")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print("Lista de Productos: ")
		for rows.Next() {
			var p Producto
			err := rows.Scan(&p.idProducto, &p.Nombre, &p.Descripcion)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(p)
		}

	}
}
