package user

import (
	"fmt"

	rest_err "github.com/15DanBerg/go-sell/config/err"
	request "github.com/15DanBerg/go-sell/model/request/user"
	"github.com/gin-gonic/gin"
)

func AddingUser(c *gin.Context) {
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		rest_err := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error=%s\n", err.Error()))

		c.JSON(rest_err.Code, rest_err)
		return
	}

	fmt.Println(userRequest)
}
