package blogs

import (
	"time"
)

// 数据库设计与模型定义
// 设计数据库表结构，至少包含以下几个表：
// users 表：存储用户信息，包括 id 、 username 、 password 、 email 等字段。
// posts 表：存储博客文章信息，包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段。
// comments 表：存储文章评论信息，包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。
// 使用 GORM 定义对应的 Go 模型结构体。

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"unique; not null" binding:"required"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null" binding:"required"`
}

type UserLogin struct { 
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type Post struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `binding:"required"`
	Content   string `binding:"required"`
	UserID    int
	User      User      `json:"omitempty"`
	Comments  []Comment `json:"omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Comment struct {
	ID        int    `gorm:"primaryKey"`
	Content   string `binding:"required"`
	UserID    int
	User      User `json:"omitempty"`
	PostID    int
	Post      Post `json:"omitempty"`
	CreatedAt time.Time
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
