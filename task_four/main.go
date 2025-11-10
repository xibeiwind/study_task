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

	r := gin.Default()
	r.Use(DBMiddleware(db))
	r.Use(AuthMiddleware())
	// r.Use(JWTMiddleware())

	r.POST("/register", blogs.Register)
	r.POST("/login", blogs.Login)

	r.GET("/articles", blogs.GetArticles)
	r.GET("/articles/:id",blogs.GetArticle)
	r.POST("/articles",  JWTMiddleware(), blogs.CreateArticle)
	r.PUT("/articles/:id",  JWTMiddleware(), blogs.UpdateArticle)
	r.DELETE("/articles/:id", JWTMiddleware(),  blogs.DeleteArticle)	

	r.GET("/comments/:post_id", blogs.GetCommentsByPostID)
	r.POST("/comments/:post_id", JWTMiddleware(),  blogs.CreateComment)
	r.DELETE("/comments/:comment_id", JWTMiddleware(),  blogs.DeleteComment)

	r.Run(":5055")

}
