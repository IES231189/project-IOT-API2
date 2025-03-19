package controllers

import (
	"api/src/StrongBox/application"
	"api/src/StrongBox/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
)


func ObtenerTodasLasCajasFuertesHandler(c *gin.Context) {
	
	repo := infraestructure.NewMongoStrongBoxRepository() 
	obtenerCajasFuertesUC := application.NewGetAllStrongBox(repo) 

	// Ejecutar la lógica de negocio para obtener todas las cajas fuertes
	strongBoxes, err := obtenerCajasFuertesUC.Execute()
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	
	c.JSON(http.StatusOK, gin.H{
		"message": "Cajas fuertes obtenidas exitosamente",
		"strongBoxes": strongBoxes,
	})
}
