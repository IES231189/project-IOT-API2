package routes

import (
	"api/src/User/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configura las rutas para los usuarios
func SetupRoutes(r *gin.Engine) {
	// Rutas para la gestión de usuarios
	r.POST("/users", controllers.CrearUsuarioHandler)                        // Crear un nuevo usuario
	r.POST("/users/:userID/invitados", controllers.AddGuestHandler)       // Agregar un invitado a un usuario
	r.GET("/users", controllers.GetAllUsersHandler)                       // Obtener todos los usuarios
	r.GET("/users/pin", controllers.GetUserByPinHandler)                  // Obtener un usuario por PIN
	r.DELETE("/users", controllers.DeleteUserHandler)                     // Eliminar un usuario
	r.DELETE("/users/:userID/invitados/:guestID", controllers.RemoveGuestHandler)
	
	r.POST("/login", controllers.LoginUserHandler)
}
