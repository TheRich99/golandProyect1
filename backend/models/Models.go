package models

import (
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nombre   string `json:"nombre"`
	Correo string `json:"correo"`
	Contraseña   string `json:"contraseña"`
}

type Restaurante struct {
	gorm.Model
	Nombre   string `json:"nombre"`
	Telefono string `json:"telefono"`
}

type Pedido struct {
	gorm.Model
	Fecha  string  `json:"fecha"`
	Total  float32 `json:"total"`
	NoMesa string  `json:"noMesa"`
}
type Plato struct {
	gorm.Model
	Nombre      string  `json:"nombre"`
	Precio      float32 `json:"precio"`
	Descripcion string  `json:"descripcion"`
	Imagen      string  `json:"imagen"`
	Stock       int     `json:"stock"`
}
type JsonPlato struct {
	Message string
	Platos  []Plato
}
type JsonUsuarios struct {
	Message string
	Usuarios  []Usuario
}

type PedidoPlato struct {
	gorm.Model
	PedidoId int    `json:"pedido_id"`
	PlatoId  int    `json:"plato_id"`
	Pedido   Pedido `gorm:"foreignKey:PedidoId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
	Plato    Plato  `gorm:"foreignKey:PlatoId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL; not null"`
}
