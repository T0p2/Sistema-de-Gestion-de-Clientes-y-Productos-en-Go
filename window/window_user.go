package window

import (
	"fmt"

	database "main/modules/DataBase"
	query "main/modules/querys"
)

func Window_user() {
	db := database.Prueba_con_bd()

	var list = query.Return_id_product(db)

	fmt.Print(list)
}
