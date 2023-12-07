package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var data = []string{"cliente1", "string", "list"}

func Window_product() {
	myApp := app.New()
	myWindow := myApp.NewWindow("List Widget")

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	myWindow.SetContent(list)
	myWindow.ShowAndRun()

	/*
	   	title_new_product,
	   	input_new_product_name,
	   	input_new_product_description, // Nuevo campo Descripción
	   	widget.NewButton("Guardar", func() {
	   		// Acción para agregar nuevo producto
	   		name := input_new_product_name.Text
	   		description := input_new_product_description.Text

	   		// Agrega aquí la lógica para guardar el nuevo producto en la base de datos
	   		query.Load_product(db, name, description)
	   	}),

	   	title_update_product,
	   	input_update_old_name_p,
	   	input_update_new_name_p,
	   	input_update_description, // Nuevo campo Nueva Descripción
	   	widget.NewButton("Actualizar", func() {
	   		// Acción para actualizar producto
	   		oldName := input_update_old_name_p.Text
	   		newName := input_update_new_name_p.Text
	   		newDescription := input_update_description.Text

	   		// Agrega aquí la lógica para actualizar el producto en la base de datos
	   		query.Update_product(db, oldName, newName, newDescription)
	   	}),

	   	title_delete_product,
	   	input_delete_product_name,
	   	widget.NewButton("Eliminar", func() {
	   		// Acción para eliminar producto
	   		productNameToDelete := input_delete_product_name.Text

	   		// Agrega aquí la lógica para eliminar el producto de la base de datos
	   		query.Delete_product(db, productNameToDelete)
	   	}),

	   )
	*/
}
