package v1

import (
	"NewGinBlog/MiddleWare"
	"NewGinBlog/Model"
	"NewGinBlog/Utills/ErrMsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login godoc
// @Summary		登录接口
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /Login   [post]
func Login(c *gin.Context) {
	var data Model.User
	c.ShouldBind(&data)
	var token string
	var code int
	code = Model.CheckLogin(data.Username, data.Password)
	if code == ErrMsg.SUCCESS {
		token, code = MiddleWare.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   ErrMsg.GetErrMessage(code),
		"token": token,
	})
}
