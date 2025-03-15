package controllers

import (
	"api/src/User/application"
	"api/src/User/infraestructure"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetUserByPinHandler maneja la obtención de un usuario por su PIN
func GetUserByPinHandler(c *gin.Context) {
	log.Println("Recibiendo solicitud GET para obtener un usuario por PIN")

	// Obtener el PIN del usuario desde la URL
	pin := c.DefaultQuery("pin", "")
	if pin == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PIN del usuario es requerido"})
		return
	}

	// Inicializar repositorio
	repo := infraestructure.NewMongoUserRepository()
	if repo == nil {
		log.Println("Error al inicializar el repositorio")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	// Crear caso de uso para obtener el usuario por PIN
	useCase := application.NewObtenerUsuarioPorPin(repo)

	// Obtener el usuario
	user, err := useCase.Ejecutar(pin)
	if err != nil {
		log.Printf("Error al obtener el usuario: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
		return
	}

	// Si no se encuentra el usuario
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Enviar la respuesta con el usuario
	c.JSON(http.StatusOK, user)
}
