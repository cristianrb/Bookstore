package http

import (
	"cristianrb/src/domain/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := c.Param("access_token_id")
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}