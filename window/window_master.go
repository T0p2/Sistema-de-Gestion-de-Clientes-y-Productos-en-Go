package window

import (
	"fmt"
	database "main/modules/DataBase"
	query "main/modules/querys"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Window_master() {
	fmt.Print("Hola mundo")

	db := database.Prueba_con_bd()

	var data_p = query.Return_id_product(db)
	var data_c = query.Return_id_client(db)

	myApp := app.New()
	myWindow_master := myApp.NewWindow("Window Master")

	var productWidgets []fyne.CanvasObject

	for _, row := range data_p {
		productID, productName := row[0], row[1]
		label := widget.NewLabel(fmt.Sprintf("Product ID: %s, Name: %s", productID, productName))
		productWidgets = append(productWidgets, label)
	}

	var clientWidget []fyne.CanvasObject

	for _, row := range data_c {
		clientID, clientName := row[0], row[1]
		label := widget.NewLabel(fmt.Sprintf("Client ID: %s, Name: %s", clientID, clientName))
		clientWidget = append(clientWidget, label)
	}

	title_new_user := widget.NewLabel("Alta de cliente")
	input_new_user_name := widget.NewEntry()
	input_new_user_name.SetPlaceHolder("Name")
	input_new_user_last_name := widget.NewEntry()
	input_new_user_last_name.SetPlaceHolder("Last Name")

	title_update_user := widget.NewLabel("Actualizar Cliente")
	input_update_old_name := widget.NewEntry()
	input_update_old_name.SetPlaceHolder("Old Name")
	input_update_new_name := widget.NewEntry()
	input_update_new_name.SetPlaceHolder("New Name")
	input_update_new_last_name := widget.NewEntry()
	input_update_new_last_name.SetPlaceHolder("New Last Name")

	title_delete_user := widget.NewLabel("Eliminar Cliente")
	input_delete_user_name := widget.NewEntry()
	input_delete_user_name.SetPlaceHolder("Name")

	title_product_user := widget.NewLabel("Carga de prodcuto a cliente, se necesita el nombre o id y el id del producto")
	title_new_product := widget.NewLabel("Agregar Producto a Cliente")
	input_new_product := widget.NewEntry()
	input_new_product.SetPlaceHolder("id Product")
	title_user_product := widget.NewLabel("Name the client")
	input_user_product := widget.NewEntry()
	input_user_product.SetPlaceHolder("name")

	content := container.NewVBox(
		title_new_user,
		input_new_user_name,
		input_new_user_last_name, // Nuevo campo Last Name
		widget.NewButton("Save", func() {
			// Acción para guardar nuevo usuario
			name := input_new_user_name.Text
			last_name := input_new_user_last_name.Text

			// Agrega aquí la lógica para guardar el nuevo usuario en la base de datos
			query.Load_user(db, name, last_name)

		}),

		title_update_user,
		input_update_old_name,
		input_update_new_name,
		input_update_new_last_name, // Nuevo campo New Last Name
		widget.NewButton("Save", func() {
			// Acción para actualizar usuario
			oldName := input_update_old_name.Text
			newName := input_update_new_name.Text
			newLastName := input_update_new_last_name.Text

			// Agrega aquí la lógica para actualizar el usuario en la base de datos
			query.Update_user(db, oldName, newName, newLastName)
		}),

		title_delete_user,
		input_delete_user_name,
		widget.NewButton("Delete", func() {
			// Acción para eliminar usuario
			userNameToDelete := input_delete_user_name.Text

			// Agrega aquí la lógica para eliminar el usuario de la base de datos
			query.Delete_user(db, userNameToDelete)
		}),

		title_product_user,
		title_user_product,
		input_user_product,
		title_new_product,
		input_new_product,
		widget.NewButton("Upload", func() {
			idProductStr := input_new_product.Text
			idProduct, err := strconv.Atoi(idProductStr)
			if err != nil {
				fmt.Print("Error")
			}

			name := input_user_product.Text

			query.New_product_user(db, idProduct, name)
		}),

		container.NewHBox(productWidgets...),
		container.NewHBox(clientWidget...),
	)

	myWindow_master.SetContent(content)
	myWindow_master.SetFullScreen(true)
	myWindow_master.ShowAndRun()
}
