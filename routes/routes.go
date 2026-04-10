
package routes

import (
    "github.com/gin-gonic/gin"
    "hris/handlers"
    "hris/middleware"
)

func SetupRoutes(r *gin.Engine) {

    tenant := r.Group("/")
    tenant.Use(middleware.TenantMiddleware())

    tenant.GET("/employees", handlers.GetEmployees)
    tenant.POST("/employees", handlers.CreateEmployee)
    tenant.PUT("/employees/:id", handlers.UpdateEmployee)
    tenant.DELETE("/employees/:id", handlers.DeleteEmployee)
}
