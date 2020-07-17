package polo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	polo = "polo"
)

func Polo(g *gin.Context) {
	g.String(http.StatusOK, polo)
}
