package handler

import (
	"net/http"

	"devyunmu.com/internal/model"

	"github.com/gin-gonic/gin"
)


func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":       "Dev云沐 — 软件工程师 · 云架构师",
		"description": model.Info.Description,
		"nav":         model.NavItems,
		"info":        model.Info,
		"skills":      model.Skills,
	})
}

func Health(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"title":       "关于 — Dev云沐",
		"description": "关于云沐 — 软件工程师、云架构师，专注于 Go 后端开发和云原生技术。",
		"nav":         model.NavItems,
		"info":        model.Info,
	})
}