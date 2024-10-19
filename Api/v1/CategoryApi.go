package v1

import (
	"NewGinBlog/Model"
	"NewGinBlog/Utills/ErrMsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加分类

// AddCategory godoc
// @Summary		用来添加分类信息
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /Cate/add  [post]
func AddCategory(c *gin.Context) {
	var data Model.Category
	_ = c.ShouldBind(&data)
	code = Model.CheckCategory(data.CategoryName)
	if code == ErrMsg.SUCCESS {
		Model.CreateCategory(&data)
	}
	if code == ErrMsg.ERROR_CATENAME_USED {
		code = ErrMsg.ERROR_CATENAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  ErrMsg.GetErrMessage(code),
		"data": data,
	})
}

//查询分类下的所有文章

// 查询分类列表

// GetCategory godoc
// @Summary		用来获取分类列表
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /Cate  [get]
func GetCategory(c *gin.Context) {
	var data []Model.Category
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code = Model.GetCategory(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  ErrMsg.GetErrMessage(code),
		"data": data,
	})
}

// 编辑分类

// EditCategory godoc
// @Summary		用来编辑分类信息
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /Cate/edit/:id   [put]
func EditCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data Model.Category
	c.ShouldBind(&data)
	code = Model.CheckCategory(data.CategoryName)
	if code == ErrMsg.SUCCESS {
		Model.EditCategory(id, &data)
	}
	if code == ErrMsg.ERROR_CATENAME_USED {
		code = ErrMsg.ERROR_CATENAME_USED
		c.Abort()
	}
	c.JSON(
		http.StatusOK, gin.H{
			"code": code,
			"msg":  ErrMsg.GetErrMessage(code),
		})
}

// 删除分类

// DeleteCategory godoc
// @Summary		用来删除分类信息
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /Cate/delete/:id  [Delete]
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := Model.DeleteCategory(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": ErrMsg.GetErrMessage(code),
		},
	)
}
