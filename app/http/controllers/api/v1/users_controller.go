// fix todo
package v1

import (
	"gohub/pkg/auth"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// Current当前用户登录信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}
