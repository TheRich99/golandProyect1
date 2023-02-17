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

<<<<<<< HEAD
	r.HandleFunc("/Login/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("./public/login.html")
		if err != nil {
			panic(err)
		}
		w.Write(file)
	})

	r.HandleFunc("/Registro/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("./public/registro.html")
		if err != nil {
			panic(err)
		}
		w.Write(file)
	})

	r.HandleFunc("/Info/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("./public/Info.html")
		if err != nil {
			panic(err)
		}
		w.Write(file)
	})

	r.HandleFunc("/miMenu/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("./public/miMenu.html")
		if err != nil {
			panic(err)
		}
		w.Write(file)
	})

	r.HandleFunc("/totalPedidos/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("./public/totalPedidos.html")
=======
	r.HandleFunc("/GestionPlatos/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("./public/Admin/gestionPlatos.html")
>>>>>>> a6609e2a42f116812791417ba7bae4e33b072f93
		if err != nil {
			panic(err)
		}
		w.Write(file)
	})

	fmt.Println("corriendo")
	http.ListenAndServe(":4000", r)
}
