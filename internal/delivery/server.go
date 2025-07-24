package delivery

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"imageflow/internal/config"
	"imageflow/internal/usecase"
	"log/slog"
	"net/http"
	"time"
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

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
