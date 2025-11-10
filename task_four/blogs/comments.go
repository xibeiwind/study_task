package blogs

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 评论功能
// 实现评论的创建功能，已认证的用户可以对文章发表评论。
func CreateComment(c *gin.Context) {
	db := getDB(c)
	if userId, exists := c.Get("user_id"); exists {
		postId, _ := strconv.Atoi(c.Param("post_id"))
		content := c.PostForm("content")
		err := db.Create(&Comment{UserID: userId.(int), PostID: postId, Content: content}).Error
		if err == nil {
			c.JSON(http.StatusOK, Response{Code: 200, Msg: "评论成功"})
		} else {
			handleError(c.Writer, err)
			c.JSON(http.StatusInternalServerError, Response{Code: 500, Msg: err.Error()})
		}
	} else {
		c.JSON(http.StatusUnauthorized, Response{Code: 500, Msg: "用户不存在"})
	}
}

// 实现评论的读取功能，支持获取某篇文章的所有评论列表。
func GetCommentsByPostID(c *gin.Context) {
	// postID int, db *gorm.DB
	db := getDB(c)
	postID := c.Param("post_id")

	var comments []Comment
	err := db.Where("post_id = ?", postID).Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusOK, Response{Code: 500, Msg: err.Error()})
	} else {
		c.JSON(http.StatusOK, Response{Code: 200, Data: comments})
	}
}
