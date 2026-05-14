package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"compress/gzip"
)


// HTTPError represents a structured HTTP error
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Code, e.Message)
}

// ErrResponse is the standard error response format
type ErrResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type gzipWriter struct {
	gin.ResponseWriter
	writer *gzip.Writer
}

func (g *gzipWriter) Write(data []byte) (int, error) {
	return g.writer.Write(data)
}

func (g *gzipWriter) WriteString(s string) (int, error) {
	return g.writer.Write([]byte(s))
}

func Gzip() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
			c.Next()
			return
		}
		gw := gzip.NewWriter(c.Writer)
		defer gw.Close()
		c.Header("Content-Encoding", "gzip")
		c.Header("Vary", "Accept-Encoding")
		c.Writer = &gzipWriter{c.Writer, gw}
		c.Next()
	}
}

func Cache() gin.HandlerFunc {
	return func(c *gin.Context) {
		maxAge := 3600
		if strings.HasPrefix(c.Request.URL.Path, "/static") {
			maxAge = 31536000
		}
		c.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", maxAge))
		c.Next()
	}
}

func Security() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		c.Header("Permissions-Policy", "camera=(), microphone=(), geolocation=()")
		c.Header("Content-Security-Policy",
			"default-src 'self'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; script-src 'self'")
		c.Next()
	}
}

// ErrorHandler is a centralized error handler middleware
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		// Log errors
		for _, err := range c.Errors {
			fmt.Printf("[%s] %s: %s\n", time.Now().Format(time.RFC3339), c.Request.URL.Path, err.Error())
		}

		// If the response hasn't been written yet, send error JSON
		if !c.Writer.Written() {
			he := &HTTPError{Code: 500, Message: "Internal Server Error"}
			if len(c.Errors) > 0 {
				if apiErr, ok := c.Errors.Last().Err.(*HTTPError); ok {
					he = apiErr
				}
			}
			c.JSON(he.Code, ErrResponse{
				Error:   http.StatusText(he.Code),
				Message: he.Message,
			})
		}
	}
}