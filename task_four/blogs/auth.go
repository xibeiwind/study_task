package blogs

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Claims 结构体（根据实际 JWT payload 自定义）
type Claims struct {
	UserID   int   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 用户认证与授权
// 实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。
// 使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。

func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPassword), err
}
func generateToken(user User, jwtSecret []byte) jwt.Token {
	// 使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。
	// 生成 JWT
	claims := &Claims{
		UserID:   int(user.ID),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "your-app",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return *token
}

// var db *gorm.DB = gorm.Open(sqlite.Open("task4.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

func getDB(c *gin.Context) *gorm.DB {
	_db, exists := c.Get("db")
	if !exists {
		return nil
	}
	db := _db.(*gorm.DB)
	return db
}

func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断user.Username是否为空
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名不能为空"})
		return
	}
	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码不能为空"})
		return
	}
	// 判断邮箱是否为空
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱不能为空"})
		return
	}

	db := getDB(c)

	if db.Where("username = ?", user.Username).First(&user).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户已存在"})
		return
	}
	if db.Where("email = ?", user.Email).First(&user).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱已存在"})
		return
	}
	if pwd, err := encryptPassword(user.Password); err != nil {
		// return err
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		user.Password = pwd
	}

	if err := db.Create(&user).Error; err != nil {
		// return err
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func Login(c *gin.Context) {
	var vm UserLogin
	if err := c.ShouldBindJSON(&vm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := getDB(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败"})
		return
	}
	var user User
	if db.Where("username = ?", vm.Username).First(&user).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户不存在"})
		return
	}
	// 验证密码
	// if pwd, err := encryptPassword(vm.Password); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// } else

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(vm.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码错误"})
		return
	} else {

		tmp, _ := c.Get("jwtSecret")
		jwtSecret := tmp.([]byte)
		// 生成 JWT
		token := generateToken(user, jwtSecret)
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "登录成功", "token": tokenString})
	}
}
