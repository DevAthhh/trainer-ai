package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/DevAthhh/trainer-ai/client/pkg/initializers"
	"github.com/DevAthhh/trainer-ai/client/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/")
}

func LoginPost(c *gin.Context) {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	data.Email = c.PostForm("email")
	data.Password = c.PostForm("password")

	var user models.UserTrainers
	initializers.DB.First(&user, "email = ?", data.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 14).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.Set("user", user.Username)
	c.Redirect(http.StatusFound, "/")
}

func RegisterPost(c *gin.Context) {
	var data struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	data.Email = c.PostForm("email")
	data.Password = c.PostForm("password")
	data.Username = c.PostForm("username")

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "hashing error",
		})
		return
	}
	user := models.UserTrainers{
		Email:    data.Email,
		Username: data.Username,
		Password: string(hash),
	}
	if result := initializers.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "db-create error",
		})
		return
	}

	c.Redirect(http.StatusFound, "/login")
}

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
