package v1

import (
	"NewGinBlog/Model"
	"NewGinBlog/Utills/ErrMsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加文章

// AddArticle godoc
// @Summary		用来添加文章信息
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /Article/add   [post]
func AddArticle(c *gin.Context) {
	var data Model.Article
	_ = c.ShouldBind(&data)
	code = Model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  ErrMsg.GetErrMessage(code),
		"data": data,
	})
}

//todo 查询分类下的所有文章
//todo 查询当个文章信息

// 查询文章列表

// GetArticle godoc
// @Summary		用来获取文章列表篇信息
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /Article  [get]
func GetArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize")) //Query返回的是一个字符串，而入参时需要入整数，需要转化
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1 //gorm提供了一个方法，如果给limited传入的为-1就没有分页功能
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := Model.GetArticle(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": ErrMsg.GetErrMessage(code),
	})
}

// GetCateArt godoc
// @Summary		用来查询分类下的文章
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /Article/cate_list/:id  [get]
func GetCateArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize")) //Query返回的是一个字符串，而入参时需要入整数，需要转化
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1 //gorm提供了一个方法，如果给limited传入的为-1就没有分页功能
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := Model.GetCateArticle(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": ErrMsg.GetErrMessage(code),
	})
}

// 编辑文章

// EditArticle godoc
// @Summary		用来编辑文章信息
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /Article/edit/:id  [put]
func EditArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data Model.Article
	c.ShouldBind(&data)
	code = Model.EditArticle(id, &data)
	c.JSON(
		http.StatusOK, gin.H{
			"code": code,
			"msg":  ErrMsg.GetErrMessage(code),
		})
}

// 删除文章

// DeleteArticle godoc
// @Summary		用来删除文章信息
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /Article/delete/:id  [delete]
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := Model.DeleteArticle(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": ErrMsg.GetErrMessage(code),
		},
	)
}

// 查询单个文章信息

// GetArtInfo godoc
// @Summary		用来查询单个文章信息
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /Article/info/{id}  [get]
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := Model.GetArtInfro(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": ErrMsg.GetErrMessage(code),
	})
}
