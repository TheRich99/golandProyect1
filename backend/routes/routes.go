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
		var platillos []models.Plato
		db.DB.Find((&platillos))

		var platos models.JsonPlato

		platos.Message = "Success"

		platos.Platos = platillos
		json.NewEncoder(w).Encode(&platos)

	case "POST":

		fmt.Println(r.FormValue("campoNombre"))
		fmt.Println(r.FormValue("precio"))
		fmt.Println(r.FormValue("cantidad"))
		fmt.Println(r.FormValue("descrip"))

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

		//w.Write([]byte(platillo.Nombre))
	}

}
