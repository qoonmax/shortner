package url

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//TODO: Create a new URL
//TODO: Get a URL

type Service interface {
	GetURL(alias string) (string, error)
	SaveURL(url string, alias string) error
}

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetURL(ctx *gin.Context) {
	alias := ctx.Param("slug")
	url, err := h.Service.GetURL(alias)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"status": "not found"})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, url)
}

func (h *Handler) SaveURL(ctx *gin.Context) {
	err := h.Service.SaveURL("url", "alias")
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}
