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

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "imageflow/docs"
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

// @Summary Upload image task
// @Description Submit an image processing task with filename
// @Tags images
// @Accept  json
// @Produce  plain
// @Param filename query string true "Image filename"
// @Success 200 {string} string "task submitted: id"
// @Failure 400 {string} string "filename required"
// @Router /upload [post]
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

// @Summary Get image task status
// @Description Get the status of image processing task by ID
// @Tags images
// @Produce json
// @Param id query string true "Task ID"
// @Success 200 {object} model.ImageTask
// @Failure 400 {string} string "id required"
// @Failure 404 {string} string "not found"
// @Router /status [get]
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
