package controllers

import (
	"api/src/User/application"
	"api/src/User/infraestructure"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func GetUserByPinHandler(c *gin.Context) {
	
	log.Println("Recibiendo solicitud GET para obtener un usuario por PIN")

	pin := c.DefaultQuery("pin", "")
	if len(pin) < 8 { 
		c.JSON(http.StatusBadRequest, gin.H{"error": "PIN del usuario es inválido"})
		return
	}

	repo := infraestructure.NewMongoUserRepository()
	if repo == nil {
		log.Println("Error al inicializar el repositorio")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	useCase := application.NewObtenerUsuarioPorPin(repo)

	user, err := useCase.Ejecutar(pin)
	if err != nil {
		log.Printf("Error al obtener el usuario: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
}
