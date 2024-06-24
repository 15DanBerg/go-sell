package product

import (
	"fmt"

	"github.com/15DanBerg/go-sell/config/validation"
	request "github.com/15DanBerg/go-sell/model/request/product"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var productRequest request.ProductRequest

	err := c.ShouldBindJSON(&productRequest)
	if err != nil {
		rest_err := validation.ValidateError(err)

		c.JSON(rest_err.Code, rest_err)
		return
	}
	fmt.Println(productRequest)
}
