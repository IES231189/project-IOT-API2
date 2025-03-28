package controllers

import (
	"api/src/User/application"
	"api/src/User/infraestructure"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoginUserHandler(c *gin.Context) {
    log.Println("Recibiendo solicitud POST para iniciar sesión")

    // Estructura para los datos de inicio de sesión
    var loginData struct {
        Correo     string `json:"correo"`
        Contraseña string `json:"contraseña"`
    }

    // Validar los datos recibidos en el cuerpo de la solicitud
    if err := c.ShouldBindJSON(&loginData); err != nil {
        log.Println("Error en los datos de entrada:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    log.Println("Correo recibido en login:", loginData.Correo) // Imprimir el correo recibido

    // Inicializar el repositorio de usuarios
    repo := infraestructure.NewMongoUserRepository()
    if repo == nil {
        log.Println("Error al inicializar el repositorio")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
        return
    }

    // Crear el servicio de inicio de sesión
    loginService := application.NewLoginUserService(repo)

    // Llamar al servicio de inicio de sesión para obtener el token y el ID de usuario
    token, userID, err := loginService.LoginUser(loginData.Correo, loginData.Contraseña)
    if err != nil {
        log.Printf("Error al iniciar sesión: %v", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
        return
    }

    // Devolver el token y el ID del usuario en la respuesta JSON
    c.JSON(http.StatusOK, gin.H{
        "message": "Inicio de sesión exitoso",
        "userID": userID,
        "token":  token,
    })
}
