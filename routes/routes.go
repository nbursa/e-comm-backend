package routes

import (
	"e-comm-backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, productController *controllers.ProductController) {
    r.OPTIONS("/*path", func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
        c.Status(http.StatusNoContent)
    })

		api := r.Group("/api")
    {
        api.GET("/products/categories", productController.GetCategories)
        api.GET("/products/category/:category", productController.GetProductsByCategory)
        api.GET("/products", productController.GetProducts)
        api.GET("/products/:id", productController.GetProduct)
        api.POST("/products", productController.CreateProduct)
        api.PUT("/products/:id", productController.UpdateProduct)
        api.DELETE("/products/:id", productController.DeleteProduct)
    }

		r.NoRoute(func(c *gin.Context) {
			c.File("./static/index.html")
		})
}