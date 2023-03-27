package handler

import (
	"net/http"

	"github.com/OtavioPontes/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show all openings
// @Description Show all openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
	}

	sendSuccess(ctx, "list-openings", openings)
}
