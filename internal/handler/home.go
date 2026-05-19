package handler

import (
	"net/http"

	"devyunmu.com/internal/model"

	"github.com/gin-gonic/gin"
)


func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":       "Dev云沐 — 全栈开发工程师 / 云原生后端架构师 · 深度交付 · 效能革新",
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
		"description": "关于云沐 — 全栈开发工程师 / 云原生后端架构师，深耕 Spring Cloud 微服务架构，系统性整合 Cursor + OpenCode 构建 AI 辅助开发工作流，具备全栈开发与性能调优能力。",
		"nav":         model.NavItems,
		"info":        model.Info,
	})
}

func Cases(c *gin.Context) {
	c.HTML(http.StatusOK, "cases.html", gin.H{
		"title":       "案例实战 — Dev云沐",
		"description": "全栈开发工程师 / 云原生后端架构师的技术案例实战",
		"nav":         model.NavItems,
		"info":        model.Info,
		"cases":       model.CaseStudies,
		"challenges":  model.DesignChallenges,
	})
}

func CaseDetail(c *gin.Context) {
	slug := c.Param("slug")
	var caseStudy *model.CaseStudy
	for i := range model.CaseStudies {
		if model.CaseStudies[i].Slug == slug {
			caseStudy = &model.CaseStudies[i]
			break
		}
	}
	if caseStudy == nil {
		var designChallenge *model.DesignChallenge
		for i := range model.DesignChallenges {
			if model.DesignChallenges[i].Slug == slug {
				designChallenge = &model.DesignChallenges[i]
				break
			}
		}
		if designChallenge == nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.HTML(http.StatusOK, "case-detail.html", gin.H{
			"title":       designChallenge.Title + " — Dev云沐",
			"description": designChallenge.Summary,
			"nav":         model.NavItems,
			"info":        model.Info,
			"case":        designChallenge,
		})
		return
	}
	c.HTML(http.StatusOK, "case-detail.html", gin.H{
		"title":       caseStudy.Title + " — Dev云沐",
		"description": caseStudy.Summary,
		"nav":         model.NavItems,
		"info":        model.Info,
		"case":        caseStudy,
	})
}