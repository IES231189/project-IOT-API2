package controllers

import (
	"api/src/User/application"
	"api/src/User/infraestructure"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetAllUsersHandler maneja la obtención de todos los usuarios
func GetAllUsersHandler(c *gin.Context) {
	log.Println("Recibiendo solicitud GET para obtener todos los usuarios")

	// Inicializar repositorio
	repo := infraestructure.NewMongoUserRepository()
	if repo == nil {
		log.Println("Error al inicializar el repositorio")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	// Crear caso de uso para obtener todos los usuarios
	useCase := application.NewObtenerTodosLosUsuarios(repo)

	// Obtener todos los usuarios
	users, err := useCase.Ejecutar()
	if err != nil {
		log.Printf("Error al obtener los usuarios: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios"})
		return
	}

	// Enviar la respuesta con la lista de usuarios
	c.JSON(http.StatusOK, users)
}
