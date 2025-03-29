package controllers

import (
	"api/src/StrongBox/application"
	"api/src/StrongBox/domain"
	"api/src/StrongBox/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CrearStrongBoxHandler maneja la solicitud para crear una nueva caja fuerte.
func CrearStrongBoxHandler(c *gin.Context) {
	// Inicializa el repositorio y el caso de uso directamente en el controlador
	repo := infraestructure.NewMongoStrongBoxRepository() // Repositorio MongoDB
	crearStrongBoxUC := application.NewCreateStrongBoxService(repo) // Servicio con repo

	// Decodificar el JSON del cuerpo de la solicitud
	var strongBox domain.StrongBox
	if err := c.ShouldBindJSON(&strongBox); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar JSON"})
		return
	}

	// Obtener el siguiente código de producto autoincremental
	codigoProducto, err := repo.ObtenerSiguienteCodigoProducto()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando código de producto"})
		return
	}
	strongBox.CodigoProducto = codigoProducto

	// Ejecutar la lógica de negocio para crear la caja fuerte
	strongBoxID, err := crearStrongBoxUC.Execute(&strongBox)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Responder con el ID de la caja fuerte creada
	c.JSON(http.StatusCreated, gin.H{
		"message":      "Caja fuerte creada exitosamente",
		"strongBox_id": strongBoxID,
		"codigo":       strongBox.CodigoProducto,
	})
}
