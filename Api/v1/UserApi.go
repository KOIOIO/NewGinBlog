package v1

import (
	"NewGinBlog/Model"
	"NewGinBlog/Utills/ErrMsg"
	"NewGinBlog/Utills/validater"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 用于接收执行数据库操作后的状态码
var code int

// AddUser godoc
// @Summary		用来添加用户
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /user/add [post]
func AddUser(c *gin.Context) {
	var data Model.User
	var msg string
	_ = c.ShouldBind(&data)

	msg, code = validater.Validate(&data)
	if code != ErrMsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  msg,
		})
		return
	}
	code = Model.CheckUser(data.Username)
	if code == ErrMsg.SUCCESS {
		Model.CreateUser(&data)
	}
	if code == ErrMsg.ERROR_USERNAME_USED {
		code = ErrMsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  ErrMsg.GetErrMessage(code),
		"data": data,
	})
}

//查询用户

// 查询用列表

// GetUsers godoc
// @Summary		用来获取用户列表
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /users  [get]
func GetUsers(c *gin.Context) {
	var data []Model.User
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code = Model.GetUsers(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  ErrMsg.GetErrMessage(code),
		"data": data,
	})
}

// 编辑用户

// EditUser godoc
// @Summary		用来修改用户信息
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /user/edit/:id  [put]
func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data Model.User
	c.ShouldBind(&data)
	code = Model.CheckUser(data.Username)
	if code == ErrMsg.SUCCESS {
		Model.EditUser(id, &data)
	}
	if code == ErrMsg.ERROR_USERNAME_USED {
		code = ErrMsg.ERROR_USERNAME_USED
		c.Abort()
	}
	c.JSON(
		http.StatusOK, gin.H{
			"code": code,
			"msg":  ErrMsg.GetErrMessage(code),
		})
}

// 删除用户

// DeleteUser godoc
// @Summary		用来删除用户信息
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /user/delete/:id  [delete]
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := Model.DeleteUser(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": ErrMsg.GetErrMessage(code),
		},
	)
}
