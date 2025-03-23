package controllers

import (
	"api/src/StrongBox/application"
	"api/src/StrongBox/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ObtenerStrongBoxByIDHandler maneja la solicitud para obtener una caja fuerte por ID.
func ObtenerStrongBoxByIDHandler(c *gin.Context) {
	// Obtener el ID de la caja fuerte desde los parámetros de la URL
	id := c.Param("id")

	// Inicializa el repositorio y el caso de uso directamente en el controlador
	repo := infraestructure.NewMongoStrongBoxRepository() // Mongo repo
	obtenerCajaFuerteUC := application.NewGetStrongBoxByID(repo) 

	// Ejecutar la lógica de negocio para obtener la caja fuerte
	box, err := obtenerCajaFuerteUC.Execute(id)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if box == nil {
		
		http.Error(c.Writer, "Caja fuerte no encontrada", http.StatusNotFound)
		return
	}

	
	c.JSON(http.StatusOK, gin.H{
		"message": "Caja fuerte encontrada",
		"strongBox": box,
	})
}
