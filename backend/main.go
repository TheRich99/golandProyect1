package main

import (
	"backend/db"
	"backend/models"
	"backend/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	fmt.Println("hola")

	db.DBConnection()

	db.DB.AutoMigrate(&models.Restaurante{}, models.Pedido{}, models.Plato{}, models.PedidoPlato{}, models.Usuario{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/platos/", routes.VistaPlato)
	r.HandleFunc("/platos/{id}", routes.VistaPlato)
	r.HandleFunc("/usuarios/", routes.VistaUsuario)

	handler := cors.Handler(r)
	http.ListenAndServe(":3000", handler)

}

func prueba() {

	mux := http.NewServeMux()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "Hello there!")
	})

	handler := cors.Handler(mux)
	http.ListenAndServe(":4000", handler)
}
