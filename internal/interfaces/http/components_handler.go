package http

import (
	"heating-prediction/internal/domain/models"
	"heating-prediction/internal/infrastructure/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func parseFloat(s string) float64 {
	if s == "" {
		return 0
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func (h *Handlers) ComponentStatusChange(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("component_id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid component id",
		})
		return
	}

	if err := h.Service.PublishDraft(
		c.Request.Context(),
		uint16(id),
		models.StatusPublished,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

//TODO сделать логгер и не возвращать чисто 400

// GET /api/component/:component_id
// GET /api/component/:component_id?next=true
func (h *Handlers) GetFeedByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("component_id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid component id"})
		return
	}

	next := c.Query("next") == "true"

	component, err := h.Service.GetFeedById(c.Request.Context(), uint16(id), next)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, component)
}

func (h *Handlers) CatalogPage(c *gin.Context) {
	minPower := parseFloat(c.Query("min_power"))

	componentsList, err := h.Service.GetFiltredComponentsList(c.Request.Context(), minPower)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, componentsList)
}

func (h *Handlers) GetCurrentUserDraft(c *gin.Context) {
	draft, err := h.Service.GetDraft(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, draft)
}

func (h *Handlers) DeleteComponent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("component_id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid component id"})
		return
	}

	if err := h.Service.DeleteComponent(c.Request.Context(), uint16(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handlers) LikeComponent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("component_id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid component id"})
		return
	}

	var body struct {
		Value int `json:"value"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.LikeComponent(c.Request.Context(), uint16(id), body.Value); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handlers) CreateComponent(c *gin.Context) {
	req := dto.ComponentRequest{}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	video, err := req.VideoFile.Open()
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer video.Close()

	image, err := req.ImageFile.Open()
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer image.Close()

	if err := h.Service.CreateComponent(c.Request.Context(), &req, video, image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "component created succusfully",
	})
}
