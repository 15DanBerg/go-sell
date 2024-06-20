package routes

import (
	"github.com/15DanBerg/go-sell/domain/product"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	//Product routes
	r.GET("/product/findById/:productId", product.FindProductByID)
	r.GET("/product/findAll", product.FindAllProduct)
	r.POST("/product/create", product.CreateProduct)
	r.PUT("/product/update/:productId", product.UpdateProduct)
	r.DELETE("/product/delete", product.DeleteProdoct)

	//Users Routes
}
