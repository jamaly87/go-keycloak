package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"keycloak/keycloak"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
)

func authMiddleware(kc *keycloak.Keycloak) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported tyoke type: %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		token := fields[1]

		// call keycloak API to verify token
		result, err := kc.Gocloak.RetrospectToken(context.Background(), token, kc.ClientId, kc.ClientSecret, kc.Realm)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		jwt, _, err := kc.Gocloak.DecodeAccessToken(context.Background(), token, kc.Realm)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		jwtj, _ := json.Marshal(jwt)
		fmt.Printf("token: %v\n", string(jwtj))

		//check if token isn't expired and valid
		if !*result.Active {
			err = fmt.Errorf("invalid or expired token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		}
		ctx.Next()
	}
}
