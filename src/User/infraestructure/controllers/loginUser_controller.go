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

    var loginData struct {
        Correo     string `json:"correo"`
        Contraseña string `json:"contraseña"`
    }

    if err := c.ShouldBindJSON(&loginData); err != nil {
        log.Println("Error en los datos de entrada:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    log.Println("Correo recibido en login:", loginData.Correo) // Imprimir el correo recibido

    repo := infraestructure.NewMongoUserRepository()
    loginService := application.NewLoginUserService(repo)

    token, err := loginService.LoginUser(loginData.Correo, loginData.Contraseña)
    if err != nil {
        log.Printf("Error al iniciar sesión: %v", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
