package routes

import (
	"api/src/User/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configura las rutas para los usuarios
func SetupRoutes(r *gin.Engine) {
	// Rutas para la gestión de usuarios
	r.POST("/users", controllers.CrearUserHandler) 
	r.POST("/usuers/:userID/invitados", controllers.AddGuestHandler)                       
	r.GET("/users", controllers.GetAllUsersHandler)          
	r.GET("/users/pin", controllers.GetUserByPinHandler)           
	r.DELETE("/users", controllers.DeleteUserHandler)  
	r.GET("/users/accesos-exitosos", controllers.GetAccesosExitososHandler)                    

}
