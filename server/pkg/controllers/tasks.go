package controllers

import (
	"net/http"

	"github.com/DevAthhh/trainer-ai/server/utils"
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

	prompt := "Можешь составить " + question.Count + " вопросов по специальности " + question.Specialty + " и со сложностью " + question.Difficulity + " Как на собеседовании. Ответь мне сообщением, содержащим только ответы на вопросы (номер - вопрос в кавычках)"
	c.JSON(http.StatusOK, gin.H{
		"status": utils.ChatStream(prompt, "deepseek/deepseek-r1-distill-llama-70b:free"),
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
	prompt := "Можешь проверить эти ответы " + checkin.Answers + " с этими вопросами " + checkin.Tasks + " и прислать сообщение в баллах от 0 до 10 с дробными, и через тире опиши одним словом уровень"
	c.JSON(http.StatusOK, gin.H{
		"status": utils.ChatStream(prompt, "deepseek/deepseek-r1-distill-llama-70b:free"),
	})
}
