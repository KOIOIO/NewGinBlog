package Routers

import (
	"NewGinBlog/Api/v1"
	"NewGinBlog/MiddleWare"
	"NewGinBlog/Utills"
	_ "NewGinBlog/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			gin和gorm的一个博客项目
// @version         2.0
// @description    本文档使用Swagger2.0标准编写的API文档

// @contact.name   Wenyu Wang

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth(JWT)

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func InitRouter() {
	gin.SetMode(Utills.AppMode)
	r := gin.New()
	r.Use(MiddleWare.LoggerMiddleware())
	r.Use(gin.Recovery())
	auth := r.Group("api/v1")
	auth.Use(MiddleWare.JwtToken())
	{
		//User
		auth.PUT("user/edit/:id", v1.EditUser)
		auth.DELETE("user/delete/:id", v1.DeleteUser)

		//Artical
		auth.POST("Article/add", v1.AddArticle)
		auth.PUT("Article/edit/:id", v1.EditArticle)
		auth.DELETE("Article/delete/:id", v1.DeleteArticle)

		//Category
		auth.POST("Cate/add", v1.AddCategory)
		auth.PUT("Cate/edit/:id", v1.EditCategory)
		auth.DELETE("Cate/delete/:id", v1.DeleteCategory)
	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("Article", v1.GetArticle)
		router.GET("Article/cate_list/:id", v1.GetCateArt)
		router.GET("Article/info/:id", v1.GetArtInfo)
		router.GET("Cate", v1.GetCategory)
		router.POST("Login", v1.Login)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = r.Run(Utills.HttpPort)

}
