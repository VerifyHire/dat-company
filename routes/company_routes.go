package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/go-microservice/controllers"
)

// SetupRoutes sets up all the routes for the company API ---
func SetupRoutes(r *gin.Engine) {
	r.GET("/companies", controllers.GetCompanies)
	r.POST("/companies", controllers.CreateCompany)
	r.PUT("/companies/:id", controllers.UpdateCompany)
	r.PATCH("/companies/:id", controllers.PatchCompany)
	r.DELETE("/companies/:id", controllers.DeleteCompany)
}
