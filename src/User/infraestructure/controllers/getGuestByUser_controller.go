package controllers

import (
	"api/src/User/application"
	"api/src/User/infraestructure"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetUserGuestsHandler maneja la solicitud para obtener los invitados de un usuario
func GetUserGuestsHandler(c *gin.Context) {
	log.Println("Recibiendo solicitud GET para obtener los invitados de un usuario")

	// Obtener el userID desde los parámetros de la URL
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID de usuario es inválido"})
		return
	}

	// Inicializar el repositorio de usuarios
	repo := infraestructure.NewMongoUserRepository()
	if repo == nil {
		log.Println("Error al inicializar el repositorio")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	// Crear el caso de uso para obtener los invitados por userID
	useCase := application.NewObtenerInvitadosPorID(repo)

	// Obtener los invitados del usuario
	guests, err := useCase.Ejecutar(userID)
	if err != nil {
		log.Printf("Error al obtener los invitados: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los invitados"})
		return
	}

	if len(guests) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron invitados para este usuario"})
		return
	}

	// Devolver la lista de invitados como respuesta JSON
	c.JSON(http.StatusOK, guests)
}
