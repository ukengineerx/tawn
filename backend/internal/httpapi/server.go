package httpapi

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"tawn/backend/internal/notes"
)

const maxNoteLength = 500

// New builds the HTTP server. A nil store keeps /health up and reports
// the database as down on /ready and /v1/notes.
func New(store notes.Store) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), cors())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/ready", func(c *gin.Context) {
		if store == nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unavailable", "postgres": "down"})
			return
		}
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		if err := store.Ping(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unavailable", "postgres": "down"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "postgres": "up"})
	})

	v1 := router.Group("/v1")
	v1.GET("/notes", func(c *gin.Context) {
		list, err := requireStore(c, store)
		if err != nil {
			return
		}
		notes, err := list.List(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not list notes"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"notes": notes})
	})
	v1.POST("/notes", func(c *gin.Context) {
		list, err := requireStore(c, store)
		if err != nil {
			return
		}
		var body struct {
			Body string `json:"body"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "body is required"})
			return
		}
		text := strings.TrimSpace(body.Body)
		if text == "" || len(text) > maxNoteLength {
			c.JSON(http.StatusBadRequest, gin.H{"error": "body must be 1 to 500 characters"})
			return
		}
		note, err := list.Create(c.Request.Context(), text)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save note"})
			return
		}
		c.JSON(http.StatusCreated, note)
	})

	return router
}

func requireStore(c *gin.Context, store notes.Store) (notes.Store, error) {
	if store == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "database unavailable"})
		return nil, notes.ErrUnavailable
	}
	return store, nil
}

func cors() gin.HandlerFunc {
	allowed := map[string]struct{}{
		"http://localhost:3000": {},
		"http://127.0.0.1:3000": {},
	}
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if _, ok := allowed[origin]; ok {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Vary", "Origin")
			c.Header("Access-Control-Allow-Headers", "Content-Type")
			c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		}
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
