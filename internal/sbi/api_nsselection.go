package sbi

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"

	"github.com/nycu-ucr/nssf/internal/logger"
	"github.com/nycu-ucr/nssf/internal/sbi/processor"
	"github.com/nycu-ucr/nssf/internal/util"
	"github.com/nycu-ucr/openapi/models"
)

func (s *Server) getNsSelectionRoutes() []Route {
	return []Route{
		{
			"Health Check",
			http.MethodGet,
			"/",
			func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"status": "Service Available"})
			},
		},

		{
			"NSSelectionGet",
			http.MethodGet,
			"/network-slice-information",
			s.NetworkSliceInformationGet,
		},
	}
}

func (s *Server) NetworkSliceInformationGet(c *gin.Context) {
	logger.NsselLog.Infof("Handle NSSelectionGet")

	var query processor.NetworkSliceInformationGetQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		logger.NsselLog.Errorf("BindQuery failed: %+v", err)
		problemDetail := &models.ProblemDetails{
			Title:         "Malformed Request",
			Status:        http.StatusBadRequest,
			Detail:        err.Error(),
			Instance:      "",
			InvalidParams: util.BindErrorInvalidParamsMessages(err),
		}
		util.GinProblemJson(c, problemDetail)
		return
	}

	s.Processor().NSSelectionSliceInformationGet(c, query)
}
