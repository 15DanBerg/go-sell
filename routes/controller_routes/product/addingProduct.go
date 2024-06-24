package product

import (
	"fmt"

	rest_err "github.com/15DanBerg/go-sell/config/err"
	request "github.com/15DanBerg/go-sell/model/request/product"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var productRequest request.ProductRequest

	err := c.ShouldBindJSON(&productRequest)
	if err != nil {
		rest_err := rest_err.NewBadRequestError(fmt.Sprintf("There are some icorrect fields, error=%s\n", err.Error()))

		c.JSON(rest_err.Code, rest_err)
		return
	}
	fmt.Println(productRequest)
}
