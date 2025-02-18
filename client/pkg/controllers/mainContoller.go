package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}

func Questions(c *gin.Context) {
	var question struct {
		Count       string `json:"count"`
		Specialty   string `json:"specialty"`
		Difficulity string `json:"difficulity"`
	}

	question.Count = c.PostForm("count")
	question.Specialty = c.PostForm("lang")
	question.Difficulity = c.PostForm("difficulity")

	jsons, _ := json.Marshal(question)

	resp, err := http.Post("http://localhost:8000/api/v1/request", "application/json", bytes.NewBuffer(jsons))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.HTML(http.StatusOK, "questions.html", gin.H{
		"Res": string(body),
	})
}
