package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func HelloworldHandler(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
