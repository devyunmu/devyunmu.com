package handler

import (
	"net/http"

	"devyunmu.com/internal/model"

	"github.com/gin-gonic/gin"
)


func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":       "Dev云沐 — AI 原生开发工程师 / 云原生后端架构师 · 全栈交付 · 效能革新",
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
		"description": "关于云沐 — AI 原生开发工程师 / 云原生后端架构师，深耕 Spring Cloud 微服务架构，系统性整合 Cursor + OpenCode 构建 AI 原生开发工作流，具备全栈开发与性能调优能力。",
		"nav":         model.NavItems,
		"info":        model.Info,
	})
}