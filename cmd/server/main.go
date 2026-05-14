package main

import (
	"html/template"
	"log"
	"time"

	"devyunmu.com/config"
	"devyunmu.com/internal/handler"
	"devyunmu.com/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"now": func() map[string]int {
			return map[string]int{"Year": time.Now().Year()}
		},
	})

	// Middleware: order matters
	r.Use(middleware.Gzip())
	r.Use(middleware.Security())
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.Cache())

	// Templates
	r.LoadHTMLGlob("templates/*.html")

	// Static assets
	r.Static("/static", "./static")
	r.StaticFile("/robots.txt", "./static/robots.txt")
	r.StaticFile("/sitemap.xml", "./static/sitemap.xml")
	r.StaticFile("/favicon.ico", "./static/images/icon.svg")
	r.StaticFile("/favicon.svg", "./static/images/icon.svg")

	// Routes
	r.GET("/health", handler.Health)
	r.GET("/", handler.Home)
	r.GET("/about", handler.About)

	addr := ":" + cfg.Port
	log.Printf("Starting server on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}