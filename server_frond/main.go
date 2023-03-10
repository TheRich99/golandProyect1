package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const DIREC = "localhost"

type Direccion struct {
	Name string
}

func Show(w http.ResponseWriter, r *http.Request) {

	user := Direccion{Name: DIREC} //creamos un nuevo usuario

	t, _ := template.ParseFiles("./public/Admin/gestionPlatos.html") //indicamos a ruta del template

	t.Execute(w, user) //le pasamos al usuario al template para que este disponible

}

func main() {
	direct := Direccion{Name: "localhost"}
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

	r.HandleFunc("/GestionPlatos/", func(w http.ResponseWriter, r *http.Request) {
		/* 		file, err := ioutil.ReadFile("./public/Admin/gestionPlatos.html")
		   		if err != nil {
		   			panic(err)
		   		}
		   		w.Write(file) */

		tmpl, err := template.ParseFiles("./public/Admin/gestionPlatos.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, direct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})
	r.HandleFunc("/GestionPlatos2/", Show)

	fmt.Println("corriendo")
	http.ListenAndServe(":4000", r)
}
