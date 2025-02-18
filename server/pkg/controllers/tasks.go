package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Request(c *gin.Context) {
	var question struct {
		Count       string `json:"count"`
		Specialty   string `json:"specialty"`
		Difficulity string `json:"difficulity"`
	}
	if c.Bind(&question) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "some request",
	})
}

func Check(c *gin.Context) {
	var checkin struct {
		Tasks   string `json:"tasks"`
		Answers string `json:"answers"`
	}
	if c.Bind(&checkin) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "some request",
	})
}
