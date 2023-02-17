package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("./public/index.html")
		if err != nil {
			panic(err)
		}
		w.Write(file)
	})

	r.HandleFunc("/Pedidos/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("./public/Pedido.html")
		if err != nil {
			panic(err)
		}
		w.Write(file)
	})

	r.HandleFunc("/Menu/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("./public/Menu.html")
		if err != nil {
			panic(err)
		}
		w.Write(file)
	})

	r.HandleFunc("/GestionPlatos/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("./public/Admin/gestionPlatos.html")
		if err != nil {
			panic(err)
		}
		w.Write(file)
	})

	fmt.Println("corriendo")
	http.ListenAndServe(":4000", r)
}
