package infragin

import (
	"fmt"
	"log"
	"net/http"
	"sesi-10/internal/config"
	"sesi-10/utility"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type middleware struct{}

func NewMiddleware() middleware {
	return middleware{}
}

func (m middleware) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("Authorization")

		if authorization == "" {
			log.Println("no token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "unauthorized",
			})
			return
		}

		token := strings.Split(authorization, "Bearer ")
		log.Println(token[1])

		if len(token) != 2 {
			log.Println("invalid len token with token", authorization)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "unauthorized",
			})
			return
		}

		tokenClaims, err := utility.VerifyToken(token[1], config.Cfg.App.Token)
		if err != nil {
			log.Println("error when try to VerifyToken with detail", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "unauthorized",
			})
			return
		}

		authId := tokenClaims["id"]
		role := "user"

		authIdInt, _ := strconv.Atoi(fmt.Sprintf("%v", authId))
		ctx.Set("auth_id", authIdInt)
		ctx.Set("role", role)

		// process next to another handler
		ctx.Next()
	}
}

func (m middleware) CheckRole(allowed []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isExists := false

		myRole := ctx.GetString("role")
		for _, allowedRole := range allowed {
			if myRole == allowedRole {
				isExists = true
				break
			}
		}

		if !isExists {
			log.Println("your role", myRole, "is prohibitted")
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"error":   "forbidden access",
			})
			return
		}

		ctx.Next()
	}
}
