package controllers

import (
	"api/src/User/application"
	"api/src/User/domain"
	"api/src/User/infraestructure"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CrearUsuarioHandler maneja la creación de un nuevo usuario
func CrearUsuarioHandler(c *gin.Context) {
	log.Println("Método recibido: POST")

	var user domain.User

	// Parsear el cuerpo de la solicitud JSON a la estructura del usuario
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error al parsear los datos: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Solicitud inválida"})
		return
	}

	// Crear una instancia del repositorio para el caso de uso
	repo := infraestructure.NewMongoUserRepository()

	// Crear el caso de uso para crear el usuario
	useCase := application.NewCrearUsuario(repo)

	// Llamar al caso de uso para crear el usuario (se encriptará la contraseña dentro del caso de uso)
	userID, err := useCase.Ejecutar(&user)
	if err != nil {
		log.Printf("Error al crear usuario: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el usuario"})
		return
	}

	log.Println("Usuario creado correctamente con ID:", userID)
	// Devolver una respuesta con el ID del usuario creado
	c.JSON(http.StatusOK, gin.H{"message": "Usuario creado correctamente", "user_id": userID})
}
