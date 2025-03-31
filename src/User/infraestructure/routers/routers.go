package routes

import (
	"api/src/User/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)


func SetupRoutes(r *gin.Engine) {
	
	r.POST("/users", controllers.CrearUsuarioHandler)                       
	r.POST("/users/:userID/invitados", controllers.AddGuestHandler)          
	r.GET("/users", controllers.GetAllUsersHandler)                          
	r.GET("/users/pin", controllers.GetUserByPinHandler)                     // Obtener un usuario por PIN
	r.GET("/users/:userID/invitados", controllers.GetUserGuestsHandler)      // Obtener los invitados de un usuario
	r.DELETE("/users", controllers.DeleteUserHandler)                        // Eliminar un usuario
	r.DELETE("/users/:userID/invitados/:guestID", controllers.RemoveGuestHandler) // Eliminar un invitado
	
	r.POST("/login", controllers.LoginUserHandler)
}
