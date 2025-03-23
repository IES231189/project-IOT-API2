package controllers

import (
	"api/src/StrongBox/application"
	"api/src/StrongBox/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

// EliminarStrongBoxHandler maneja la solicitud para eliminar una caja fuerte por ID.
func EliminarStrongBoxHandler(c *gin.Context) {
	// Obtener el ID de la caja fuerte desde los parámetros de la URL
	id := c.Param("id")

	// Inicializa el repositorio y el caso de uso directamente en el controlador
	repo := infraestructure.NewMongoStrongBoxRepository() // Mongo repo
	eliminarCajaFuerteUC := application.NewDeleteStrongBox(repo) // Servicio con repo

	// Ejecutar la lógica de negocio para eliminar la caja fuerte
	err := eliminarCajaFuerteUC.Execute(id)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{
		"message": "Caja fuerte eliminada exitosamente",
	})
}
