package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
// Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
type User struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Posts     []Post `gorm:"foreignKey:UserId"`
	PostCount int
}
type Post struct {
	Id            int `gorm:"primaryKey"`
	UserId        int
	Title         string
	Content       string
	Comments      []Comment `gorm:"foreignKey:PostId"`
	CommentStatus string
}
type Comment struct {
	Id      int `gorm:"primaryKey"`
	PostId  int
	Content string
	UserId  int
}

func SetupBlogTables(db *gorm.DB) {
	db.Migrator().DropTable(&User{}, &Post{}, &Comment{})

	// 迁移模型
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

func TestStep1() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	SetupBlogTables(db)
}

// 创建测试数据
func InitBlogData(db *gorm.DB) {

	// 创建多个用户和多篇post，每个post有多个用户的comment
	user1 := User{Name: "Alice"}
	db.Create(&user1)
	user2 := User{Name: "Bob"}
	db.Create(&user2)
	user3 := User{Name: "Charlie"}
	db.Create(&user3)
	post1 := Post{UserId: user1.Id, Title: "First Post", Content: "This is my first post."}	
	db.Create(&post1)
	post2 := Post{UserId: user2.Id, Title: "Second Post", Content: "This is my second post."}
	db.Create(&post2)
	comment1 := Comment{PostId: post1.Id, Content: "Nice post!", UserId: user2.Id}
	db.Create(&comment1)
	comment2 := Comment{PostId: post2.Id, Content: "Great post!", UserId: user1.Id}
	db.Create(&comment2)
	comment3 := Comment{PostId: post1.Id, Content: "I like this post!", UserId: user1.Id}
	db.Create(&comment3)
	comment4 := Comment{PostId: post1.Id, Content: "I don't like this post!", UserId: user3.Id}
	db.Create(&comment4)

}

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

func QueryUserPostsAndComments(db *gorm.DB, userId int) ([]Post, error) {
	var posts []Post
	err := db.Preload("Comments").Where("user_id = ?", userId).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
func QueryMostCommentedPost(db *gorm.DB) (Post, error) {
	var post Post
	// 使用子查询找到评论数量最多的文章ID
	var postID int
	err := db.Model(&Comment{}).
		Select("post_id").
		Group("post_id").
		Order("COUNT(*) desc").
		Limit(1).
		Pluck("post_id", &postID).Error

	if err != nil {
		return post, err
	}

	// 根据找到的文章ID查询完整的文章信息和评论
	err = db.Preload("Comments").Where("id = ?", postID).First(&post).Error
	return post, err
}

func TestStep2() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	InitBlogData(db)
	posts, err := QueryUserPostsAndComments(db, 1)
	if err != nil {
		panic(err)
	}
	for _, post := range posts {
		fmt.Println("Post:", post.Title)
		for _, comment := range post.Comments {
			fmt.Println("Comment:", comment.Content)
		}
	}

	post, err := QueryMostCommentedPost(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Most commented post:", post.Title)
}

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (p *Post) AfterCreate(db *gorm.DB) (err error) {
	// 更新用户文章数量
	var user User
	db.First(&user, p.UserId)
	user.PostCount++
	db.Save(&user)
	return
}
func (c *Comment) AfterDelete(db *gorm.DB) (err error) {
	// 检查文章的评论数量
	var post Post
	db.Preload("Comments").Where("id = ?", c.PostId).First(&post)
	if len(post.Comments) == 0 {
		post.CommentStatus = "无评论"
		db.Save(&post)
	}
	return
}

func TestStep3() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	SetupBlogTables(db)
	InitBlogData(db)
	var post Post
	db.Preload("Comments").Where("id = ?", 2).First(&post)
	db.Delete(&post.Comments[0])
	db.Preload("Comments").First(&post, 2)
	fmt.Println("Post.CommentStatus:", post.CommentStatus)

}

func main() {
	TestStep1()
	TestStep2()
	TestStep3()
}
