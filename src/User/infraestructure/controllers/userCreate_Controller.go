package controllers

import (
	"api/src/User/application"
	"api/src/User/domain"
	"api/src/User/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CrearUserHandler maneja la solicitud para crear un usuario.
func CrearUserHandler(c *gin.Context) {
	// Inicializa el repositorio y el caso de uso directamente en el controlador
	repo := infraestructure.NewMongoUserRepository()
	crearUsuarioUC := application.NewCrearUsuario(repo)

	// Decodificar el JSON del cuerpo de la solicitud
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Generar un ObjectID si no está presente
	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}

	// Ejecutar la lógica de negocio para crear un usuario
	userID, err := crearUsuarioUC.Ejecutar(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el usuario"})
		return
	}

	// Responder con el ID del usuario creado
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario creado exitosamente",
		"user_id": userID,
	})
}
