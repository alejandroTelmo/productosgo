package main

import (
	"encoding/json"
	"log"

	"os"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Producto struct {
	ID             int       `json:"id"`
	Nombre         string    `json:"nombre"`
	Precio         float64   `json:"precio"`
	Stock          int       `json:"stock"`
	Codigo         string    `json:"codigo"`
	Publicado      bool      `json:"publicado"`
	FechaDeCreacion time.Time `json:"fechaDeCreacion"`
}

func main() {
	productos, err := cargarProductos()
	if err != nil {
		log.Fatal("Error al cargar los productos:", err)
	}
	imprimirProductos(productos)
	r := gin.Default()

	r.GET("/productos", func(c *gin.Context) {
		productos, err := cargarProductos()
		if err != nil {
			c.JSON(500, gin.H{"error": "Error al cargar los productos"})
			return
		}
		c.JSON(200, productos)
	})

	r.Run()
}

func cargarProductos() ([]Producto, error) {
	data, err := os.ReadFile("productos.json")
	if err != nil {
		return nil, err
	}

	var productos []Producto
	err = json.Unmarshal(data, &productos)
	if err != nil {
		return nil, err
	}

	return productos, nil
}
func imprimirProductos(productos []Producto) {
	for _, p := range productos {
		fmt.Printf("ID: %d\n", p.ID)
		fmt.Printf("Nombre: %s\n", p.Nombre)
		fmt.Printf("Precio: %.2f\n", p.Precio)
		fmt.Printf("Stock: %d\n", p.Stock)
		fmt.Printf("Código: %s\n", p.Codigo)
		fmt.Printf("Publicado: %v\n", p.Publicado)
		fmt.Printf("Fecha de Creación: %s\n", p.FechaDeCreacion.Format(time.RFC3339))
		fmt.Println("-----------------------------")
	}
}