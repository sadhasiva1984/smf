package util_oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sadhasiva1984/openapi/models"
	smf_context "github.com/sadhasiva1984/smf/internal/context"
	"github.com/sadhasiva1984/smf/internal/logger"
)

type RouterAuthorizationCheck struct {
	serviceName   models.ServiceName
	ServingPlmnID models.PlmnId
	RoamingPlmnID models.PlmnId
}

func NewRouterAuthorizationCheck(serviceName models.ServiceName) *RouterAuthorizationCheck {
	smfContext := smf_context.GetSelf()

	var servingPlmnID, roamingPlmnID models.PlmnId

	if smfContext.Roaming != nil {
		if smfContext.Roaming.ServingPlmnID != nil {
			servingPlmnID = *smfContext.Roaming.ServingPlmnID
		}
		if smfContext.Roaming.RoamingPlmnID != nil {
			roamingPlmnID = *smfContext.Roaming.RoamingPlmnID
		}
	}

	return &RouterAuthorizationCheck{
		serviceName:   serviceName,
		ServingPlmnID: servingPlmnID,
		RoamingPlmnID: roamingPlmnID,
	}
}

func (rac *RouterAuthorizationCheck) Check(c *gin.Context, smfContext smf_context.NFContext) {
	token := c.Request.Header.Get("Authorization")
	err := smfContext.AuthorizationCheck(token, rac.serviceName, rac.ServingPlmnID, rac.RoamingPlmnID)
	if err != nil {
		logger.UtilLog.Debugf("RouterAuthorizationCheck: Check Unauthorized: %s", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	logger.UtilLog.Debugf("RouterAuthorizationCheck: Check Authorized")
}
