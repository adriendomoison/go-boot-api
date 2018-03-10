/*
	Authentication with OAuth2.0 package
*/
package oauth2

import (
	"github.com/gin-gonic/gin"
)

type RestInterface interface {
	AppAuthorize(c *gin.Context)
	AppToken(c *gin.Context)
	AppInfo(c *gin.Context)
	AppAuthCode(c *gin.Context)
	AppAuthToken(c *gin.Context)
	AppAuthPassword(c *gin.Context)
	AppAuthClientCredentials(c *gin.Context)
	AppAuthAssertion(c *gin.Context)
	AppAuthRefresh(c *gin.Context)
	AppAuthInfo(c *gin.Context)
	GetAccessTokenOwnerUserId(c *gin.Context)
}

// Component
type Component struct {
	rest RestInterface
}

// New return a new micro service instance
func New(rest RestInterface) *Component {
	return &Component{rest}
}

// Attach link the oauth micro-service with its dependencies to the system
func (component *Component) AttachPublicAPI(group *gin.RouterGroup) {
	group.GET("/authorize", component.rest.AppAuthorize)
	group.POST("/authorize", component.rest.AppAuthorize)
	group.POST("/token", component.rest.AppToken)
	group.POST("/info", component.rest.AppInfo)
	group.GET("/oauth2/code", component.rest.AppAuthCode)
	group.GET("/oauth2/token", component.rest.AppAuthToken)
	group.POST("/oauth2/password", component.rest.AppAuthPassword)
	group.GET("/oauth2/client_credentials", component.rest.AppAuthClientCredentials)
	group.GET("/oauth2/assertion", component.rest.AppAuthAssertion)
	group.GET("/oauth2/refresh", component.rest.AppAuthRefresh)
	group.GET("/oauth2/info", component.rest.AppAuthInfo)
}

// Attach link the oauth micro-service with its dependencies to the system
func (component *Component) AttachPrivateAPI(group *gin.RouterGroup) {
	group.GET("/access-token/:accessToken/get-owner", component.rest.GetAccessTokenOwnerUserId)
}