package controllers

import (
	"api/src/User/application"
	"api/src/User/domain"
	"api/src/User/infraestructure"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"net/http"
)

// CrearUserHandler maneja la solicitud para crear un usuario.
func CrearUserHandler(c *gin.Context) {
	// Inicializa el repositorio y el caso de uso directamente en el controlador
	repo := infraestructure.NewMongoUserRepository()
	crearUsuarioUC := application.NewCrearUsuario(repo)

	// Decodificar el JSON del cuerpo de la solicitud
	var user domain.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		http.Error(c.Writer, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	// Ejecutar la lógica de negocio para crear un usuario
	userID, err := crearUsuarioUC.Ejecutar(&user)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con el ID del usuario creado
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario creado exitosamente",
		"user_id": userID,
	})
}
