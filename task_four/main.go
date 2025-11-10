package main

import (
	"task_four/blogs"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectSqlistDB() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("task4.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&blogs.User{}, &blogs.Post{}, &blogs.Comment{})
	return db
}

func main() {
	db := ConnectSqlistDB()
	DBMiddleware(db)

	r := gin.Default()

	r.POST("/register", blogs.Register)
	r.POST("/login", blogs.Login)

	r.GET("/articles", blogs.GetArticles)
	r.GET("/articles/:id", blogs.GetArticle)
	r.GET("/comments/:post_id", blogs.GetCommentsByPostID)

	protected := r.Group("/api")

	protected.Use(JWTMiddleware())
	{
		r.POST("/articles", blogs.CreateArticle)
		r.PUT("/articles/:id", blogs.UpdateArticle)
		r.DELETE("/articles/:id", blogs.DeleteArticle)

		r.POST("/comments", blogs.CreateComment)
	}

}
