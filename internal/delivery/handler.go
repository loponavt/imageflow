package delivery

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"imageflow/internal/config"
	"imageflow/internal/usecase"
)

type Handler struct {
	uc  *usecase.ImageUseCase
	srv *http.Server
}

func NewHandler(cfg *config.Config, uc *usecase.ImageUseCase) *Handler {
	h := &Handler{uc: uc}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	v1 := router.Group("/api/v1")
	{
		v1.POST("/upload", h.upload)
		v1.GET("/status", h.status)
	}

	h.srv = &http.Server{
		Addr:    cfg.HTTPPort,
		Handler: router,
	}
	return h
}

func (h *Handler) Start() error {
	slog.Info("HTTP server started", "addr", h.srv.Addr)
	return h.srv.ListenAndServe()
}

func (h *Handler) Stop(ctx context.Context) error {
	return h.srv.Shutdown(ctx)
}

func (h *Handler) upload(c *gin.Context) {
	filename := c.Query("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filename required"})
		return
	}
	id, err := h.uc.Submit(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task_id": id})
}

func (h *Handler) status(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id required"})
		return
	}
	task, err := h.uc.GetStatus(id)
	if err != nil || task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}
