package product

import (
	rest_err "github.com/15DanBerg/go-sell/config/err"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	err := rest_err.NewBadRequestError("Request error")
	c.JSON(err.Code, err)
}
