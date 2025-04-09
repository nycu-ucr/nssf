package util

import (
	"github.com/nycu-ucr/gin"

	"github.com/nycu-ucr/openapi/models"
)

func GinProblemJson(c *gin.Context, problemDetails *models.ProblemDetails) {
	c.JSON(int(problemDetails.Status), problemDetails)
	c.Writer.Header().Set("Content-Type", "application/problem+json")
}
