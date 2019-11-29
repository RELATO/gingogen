package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/relato/gingogen/boilerplate/mode"
)

func init() {
	groupApi.POST("login", login)
}

func login(c *gin.Context) {
	var mdl models.AuthorizationModel
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	ip := c.ClientIP()
	data, err := mdl.Login(ip)
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
