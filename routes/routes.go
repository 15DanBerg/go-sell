package routes

import (
	"github.com/15DanBerg/go-sell/domain/product"
	"github.com/15DanBerg/go-sell/domain/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	//Product routes
	r.GET("/product/findById/:productId", product.FindProductByID)
	r.GET("/product/findAll", product.FindAllProducts)
	r.POST("/product/create", product.CreateProduct)
	r.PUT("/product/update/:productId", product.UpdateProduct)
	r.DELETE("/product/delete", product.DeleteProdoct)

	//Users Routes
	r.GET("/user/findById/:userId", user.FindUserByID)
	r.GET("/user/findAll", user.FindAllUsers)
	r.POST("/user/create", user.AddingUser)
	r.PUT("/user/update/:userId", user.UpdateUser)
	r.DELETE("/user/delete", user.DeleteUser)
}
