package blogs

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getUserFromToken(c *gin.Context, db *gorm.DB) (User, error) {
	// 通过jwt token 获取用户Id
	if userId, exists := c.Get("user_id"); exists {
		user := User{}
		db.Where("id=?", userId).First(&user)
		return user, nil
	}
	return User{}, errors.New("用户不存在")
}

// 文章管理功能
// 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
func CreateArticle(c *gin.Context) {
	db := getDB(c)

	// 如何通过jwt token 获取信息
	if user, err := getUserFromToken(c, db); err != nil {
		// return err
		c.JSON(http.StatusUnauthorized, Response{Code: 500, Msg: "用户不存在"})
	} else {
		// 创建文章
		var post Post
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = db.Create(&Post{Title: post.Title, Content: post.Content, UserID: user.ID}).Error
		if err != nil {
			c.JSON(http.StatusOK, Response{Code: 500, Msg: err.Error()})
		} else {
			c.JSON(http.StatusOK, Response{Code: 200, Msg: "创建成功"})
		}
	}
}

// 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
func GetArticles(c *gin.Context) {
	db := getDB(c)
	// 获取所有文章
	var posts []Post
	err := db.Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusOK, Response{Code: 500, Msg: err.Error()})
	} else {
		c.JSON(http.StatusOK, Response{Code: 200, Data: posts})
	}
}
func GetArticle(c *gin.Context) {
	db := getDB(c)
	id := c.Param("id")
	// 获取文章
	var post Post
	err := db.First(&post, id).Error
	// return post, err
	if err != nil {
		c.JSON(http.StatusOK, Response{Code: 500, Msg: err.Error()})
	} else {
		c.JSON(http.StatusOK, Response{Code: 200, Data: post})
	}
}

// 实现文章的更新功能，只有文章的作者才能更新自己的文章。
func UpdateArticle(c *gin.Context) {
	db := getDB(c)
	id := c.Param("id")
	// 如何通过jwt token 获取信息
	if userId, exists := c.Get("user_id"); exists {
		// 更新文章
		var post Post
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := db.Model(&Post{}).Where("id = ? AND user_id = ?", id, userId).Updates(map[string]interface{}{
			"title":post.Title,
			"content":post.Content,
		}).Error
		// .Update("title", post.Title).Update("content", post.Content).Error
		if err != nil {
			c.JSON(http.StatusOK, Response{Code: 500, Msg: err.Error()})
		} else {
			c.JSON(http.StatusOK, Response{Code: 200, Msg: "更新成功"})
		}
	} else {
		c.JSON(http.StatusUnauthorized, Response{Code: 500, Msg: "用户不存在"})
	}
}

// 实现文章的删除功能，只有文章的作者才能删除自己的文章。
func DeleteArticle(c *gin.Context) {
	db := getDB(c)
	id := c.Param("id")
	if userId, exists := c.Get("user_id"); exists {
		// 删除文章
		err := db.Where("id = ? AND user_id = ?", id, userId).Delete(&Post{}).Error
		if err != nil {
			c.JSON(http.StatusOK, Response{Code: 500, Msg: err.Error()})
		} else {
			c.JSON(http.StatusOK, Response{Code: 200, Msg: "删除成功"})
		}
	} else {
		c.JSON(http.StatusUnauthorized, Response{Code: 500, Msg: "用户不存在"})
	}
}
