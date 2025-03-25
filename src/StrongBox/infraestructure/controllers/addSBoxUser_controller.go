package controllers

import (
	"api/src/StrongBox/application"
	"api/src/StrongBox/domain"
	"api/src/StrongBox/infraestructure"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddUserToStrongBox maneja la solicitud para agregar un usuario a una caja fuerte.
func AddUserToStrongBox(c *gin.Context) {
	// ✅ Inicializa el repositorio y el caso de uso directamente en el controlador
	repo := infraestructure.NewMongoStrongBoxRepository() // Mongo repo
	addUserUC := application.NewAddUserToStrongBoxService(repo) // Servicio con repo

	// ✅ Obtener el ID de la caja fuerte desde la URL
	boxID := c.Param("boxID")

	// ✅ Decodificar el JSON del cuerpo de la solicitud
	var user domain.UserStrongBox
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		http.Error(c.Writer, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	// ✅ Ejecutar la lógica de negocio para agregar el usuario
	err = addUserUC.Execute(boxID, &user)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ Responder con éxito
	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario agregado exitosamente a la caja fuerte",
	})
}
