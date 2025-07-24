package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Upload image task
// @Description Submit an image processing task with filename
// @Tags images
// @Accept  json
// @Produce  plain
// @Param filename query string true "Image filename"
// @Param type query string true "Processing type: resize | grayscale | blur"
// @Success 200 {string} string "task submitted: id"
// @Failure 400 {string} string "filename required"
// @Router /upload [post]
func (h *Handler) upload(c *gin.Context) {
	filename := c.Query("filename")
	taskType := c.Query("type")
	if filename == "" || taskType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filename and type required"})
		return
	}
	id, err := h.uc.Submit(filename, taskType)
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
