package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type doc struct {
	Id   string    `json:"id"`
	Num  string    `json:"num"`
	Date time.Time `json:"date"`
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
}

func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	jwt, err := server.keycloak.Gocloak.Login(context.Background(),
		server.keycloak.ClientId,
		server.keycloak.ClientSecret,
		server.keycloak.Realm,
		req.Username,
		req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, errorResponse(err))
		return
	}
	resp := &loginResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) getDocs(ctx *gin.Context) {
	rs := []*doc{
		{
			Id:   "1",
			Num:  "ABC-123",
			Date: time.Now().UTC(),
		},
		{
			Id:   "2",
			Num:  "ABC-456",
			Date: time.Now().UTC(),
		},
	}

	ctx.JSON(http.StatusOK, rs)
}
