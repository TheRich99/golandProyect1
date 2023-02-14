package main

import (
	"backend/db"
	"backend/models"
	"backend/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hola")

	db.DBConnection()

	db.DB.AutoMigrate(&models.Restaurante{}, models.Pedido{}, models.Plato{}, models.PedidoPlato{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/platos/", routes.VistaPlato)
	r.HandleFunc("/platos/{id}", routes.VistaPlato)
	http.ListenAndServe(":3000", r)

}
