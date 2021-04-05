package infrastructure

import (
	"github.com/HideBa/go-cleanarch/app/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSqlHandler())
	articleController := controllers.NewArticleController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })
	router.POST("/articles", func(c *gin.Context) { articleController.Create(c) })
	router.GET("articles/:id", func(c *gin.Context) { articleController.Show(c) })
	router.GET("articles", func(c *gin.Context) { articleController.Index(c) })
	Router = router
}
