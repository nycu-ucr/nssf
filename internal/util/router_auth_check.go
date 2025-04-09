package util

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"

	nssf_context "github.com/nycu-ucr/nssf/internal/context"
	"github.com/nycu-ucr/nssf/internal/logger"
	"github.com/nycu-ucr/openapi/models"
)

type RouterAuthorizationCheck struct {
	serviceName models.ServiceName
}

func NewRouterAuthorizationCheck(serviceName models.ServiceName) *RouterAuthorizationCheck {
	return &RouterAuthorizationCheck{
		serviceName: serviceName,
	}
}

func (rac *RouterAuthorizationCheck) Check(c *gin.Context, nssfContext nssf_context.NFContext) {
	token := c.Request.Header.Get("Authorization")
	err := nssfContext.AuthorizationCheck(token, rac.serviceName)
	if err != nil {
		logger.UtilLog.Debugf("RouterAuthorizationCheck: Check Unauthorized: %s", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	logger.UtilLog.Debugf("RouterAuthorizationCheck: Check Authorized")
}
