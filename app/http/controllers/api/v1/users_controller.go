// fix todo
package v1

import (
	user "gohub/app/models/users"
	"gohub/app/requests"
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

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Paginatoin); !ok {
		return
	}

	data, page := user.Paginate(c, 5)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": page,
	})

}
