package user

import (
	"fmt"

	"github.com/15DanBerg/go-sell/config/validation"
	request "github.com/15DanBerg/go-sell/model/request/user"
	"github.com/gin-gonic/gin"
)

func AddingUser(c *gin.Context) {
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		rest_err := validation.ValidateError(err)

		c.JSON(rest_err.Code, rest_err)
		return
	}

	fmt.Println(userRequest)
}
