package controllers

import (
	"api/src/StrongBox/application"
	"api/src/StrongBox/domain"
	"api/src/StrongBox/infraestructure"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"net/http"
)

// CrearStrongBoxHandler maneja la solicitud para crear una nueva caja fuerte.
func CrearStrongBoxHandler(c *gin.Context) {
	// Inicializa el repositorio y el caso de uso directamente en el controlador
	repo := infraestructure.NewMongoStrongBoxRepository() // Mongo repo
	crearStrongBoxUC := application.NewCreateStrongBoxService(repo) // Servicio con repo

	// Decodificar el JSON del cuerpo de la solicitud
	var strongBox domain.StrongBox
	err := json.NewDecoder(c.Request.Body).Decode(&strongBox)
	if err != nil {
		http.Error(c.Writer, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	// Ejecutar la lógica de negocio para crear la caja fuerte
	strongBoxID, err := crearStrongBoxUC.Execute(&strongBox)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con el ID de la caja fuerte creada
	c.JSON(http.StatusCreated, gin.H{
		"message": "Caja fuerte creada exitosamente",
		"strongBox_id": strongBoxID,
	})
}
