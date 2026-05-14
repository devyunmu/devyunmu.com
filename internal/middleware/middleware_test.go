package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func newTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestHTTPError(t *testing.T) {
	err := &HTTPError{Code: 404, Message: "Not Found"}
	if err.Error() != "HTTP 404: Not Found" {
		t.Errorf("expected 'HTTP 404: Not Found', got %q", err.Error())
	}
}

func TestErrResponse(t *testing.T) {
	resp := ErrResponse{Error: "Not Found", Message: "Resource not found"}
	if resp.Error != "Not Found" {
		t.Errorf("expected 'Not Found', got %q", resp.Error)
	}
	if resp.Message != "Resource not found" {
		t.Errorf("expected 'Resource not found', got %q", resp.Message)
	}
}
func TestGzip(t *testing.T) {
	r := gin.New()
	r.Use(Gzip())
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	r.ServeHTTP(w, req)

	if w.Header().Get("Content-Encoding") != "gzip" {
		t.Errorf("expected gzip encoding, got %q", w.Header().Get("Content-Encoding"))
	}
	if !strings.Contains(w.Header().Get("Vary"), "Accept-Encoding") {
		t.Errorf("expected Vary: Accept-Encoding, got %q", w.Header().Get("Vary"))
	}
}

func TestGzipNoAccept(t *testing.T) {
	r := newTestRouter()
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(w, req)

	if w.Header().Get("Content-Encoding") != "" {
		t.Errorf("expected no encoding, got %q", w.Header().Get("Content-Encoding"))
	}
}

func TestCache(t *testing.T) {
	r := newTestRouter()
	mw := Cache()
	r.GET("/static/test.css", mw, func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.GET("/other", mw, func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/static/test.css", nil)
	r.ServeHTTP(w, req)

	cc := w.Header().Get("Cache-Control")
	if !strings.Contains(cc, "31536000") {
		t.Errorf("static should get 1yr cache, got: %s", cc)
	}

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest(http.MethodGet, "/other", nil)
	r.ServeHTTP(w2, req2)

	cc2 := w2.Header().Get("Cache-Control")
	if strings.Contains(cc2, "31536000") {
		t.Errorf("non-static should not get 1yr cache, got: %s", cc2)
	}
}

func TestSecurityHeaders(t *testing.T) {
	r := newTestRouter()
	r.GET("/test", Security(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(w, req)

	if w.Header().Get("X-Content-Type-Options") != "nosniff" {
		t.Errorf("expected nosniff, got %q", w.Header().Get("X-Content-Type-Options"))
	}
	if w.Header().Get("X-Frame-Options") != "DENY" {
		t.Errorf("expected DENY, got %q", w.Header().Get("X-Frame-Options"))
	}
	if !strings.Contains(w.Header().Get("Content-Security-Policy"), "default-src 'self'") {
		t.Errorf("expected CSP default-src 'self', got %q", w.Header().Get("Content-Security-Policy"))
	}
	if !strings.Contains(w.Header().Get("Strict-Transport-Security"), "63072000") {
		t.Errorf("expected HSTS, got %q", w.Header().Get("Strict-Transport-Security"))
	}
}

func TestErrorHandler(t *testing.T) {
	r := newTestRouter()
	r.Use(ErrorHandler())
	r.GET("/error", func(c *gin.Context) {
		c.Error(&HTTPError{Code: 404, Message: "Not Found"})
	})
	r.GET("/ok", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/error", nil)
	r.ServeHTTP(w, req)

	if w.Code != 404 {
		t.Errorf("expected 404, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Not Found") {
		t.Errorf("expected error body, got %q", w.Body.String())
	}

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest(http.MethodGet, "/ok", nil)
	r.ServeHTTP(w2, req2)

	if w2.Code != 200 {
		t.Errorf("expected 200, got %d", w2.Code)
	}
}
