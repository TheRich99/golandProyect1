package routes

import (
	"backend/db"
	models "backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hola"))
}
func VistaPlato(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var usuarios []models.Plato
		db.DB.Find((&usuarios))

		var platos models.JsonPlato

		platos.Message = "Success"

		platos.Platos = usuarios
		json.NewEncoder(w).Encode(&platos)

	case "POST":
		fmt.Println(r.FormValue("campoNombre"))
		fmt.Println(r.FormValue("precio"))
		fmt.Println(r.FormValue("cantidad"))
		fmt.Println(r.FormValue("descrip"))
		fmt.Println("hola")
		var platillo models.Plato

		platillo.Nombre = r.FormValue("campoNombre")
		platillo.Descripcion = r.FormValue("descrip")

		if s, err := strconv.ParseFloat(r.FormValue("precio"), 32); err == nil {
			platillo.Precio = float32(s)
		}

		platillo.Stock = 0

		//json.NewDecoder(r.Body).Decode(&platillo)

		creado := db.DB.Create(&platillo)

		if creado.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(creado.Error.Error()))
		}
		json.NewEncoder(w).Encode(&platillo)

	case "DELETE":
		var plato models.Plato
		parametros := mux.Vars(r)
		db.DB.First(&plato, parametros["id"])

		if plato.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Plato no encontrado"))
			return
		}
		db.DB.Delete(&plato)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Eliminado Correctamente"))

	case "PUT":
		var nuevo models.Plato
		var anterior models.Plato
		json.NewDecoder(r.Body).Decode(&nuevo)
		db.DB.First(&anterior, nuevo.ID)
		if anterior.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Plato no encontrado"))
			return
		} else {
			db.DB.Model(&anterior).Updates(&nuevo)
			w.Write([]byte("Plato  editado"))
		}

		//w.Write([]byte(usuarios.Nombre))
	}

}

func VistaUsuario(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var user models.Usuario
		user.Correo = r.FormValue("email")
		user.Contraseña = r.FormValue("pass")

		err := json.NewDecoder(r.Body).Decode(&user)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprintf(w, "Solicitud inválida: %v", err)
            return
        }
		//if err := db.DB.Where("correo = ? AND contraseña = ?", user.Correo, user.Contraseña).First(&user).Error; err != nil {
        if err := db.DB.Where("correo = ? AND contraseña = ?", r.FormValue("email"), r.FormValue("pass")).First(&user).Error; err != nil {
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprint(w, "Credenciales incorrectas, intente nuevamente")
            return
        }
        json.NewEncoder(w).Encode(&user)


	case "POST":
		//fmt.Println(r.FormValue("cedula"))
		//fmt.Println(r.FormValue("nombre"))
		//fmt.Println(r.FormValue("correo"))
		//fmt.Println(r.FormValue("contrasena"))
		
		var usuarios models.Usuario

		usuarios.Nombre = r.FormValue("nombre")
		usuarios.Contraseña = r.FormValue("contrasena")
		usuarios.Correo = r.FormValue("correo")

		if s, err := strconv.ParseUint(r.FormValue("cedula"), 10, 32); err == nil {
			usuarios.ID = uint(s)
		}
		
		creado := db.DB.Create(&usuarios)
		
		if creado.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(creado.Error.Error()))
		}else{
			http.Redirect(w, r, "http://localhost:4000/Registro/", http.StatusSeeOther)
		}
		json.NewEncoder(w).Encode(&usuarios)
	}
}
