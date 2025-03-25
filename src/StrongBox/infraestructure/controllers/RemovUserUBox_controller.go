package controllers

import (
	"api/src/StrongBox/application"
	"api/src/StrongBox/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RemoveUserFromStrongBoxHandler maneja la solicitud para eliminar un usuario de una caja fuerte.
func RemoveUserFromStrongBoxHandler(c *gin.Context) {
	boxID := c.Param("boxID")
	userID := c.Param("userID")

	if boxID == "" || userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faltan parámetros en la solicitud"})
		return
	}

	// Inicializamos el repositorio y el servicio
	repo := infraestructure.NewMongoStrongBoxRepository()
	removeUserService := application.NewRemoveUserFromStrongBoxService(repo)

	// Ejecutamos la lógica de eliminación
	err := removeUserService.Execute(boxID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente de la caja fuerte"})
}
