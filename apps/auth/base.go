package auth

import (
	infragin "sesi-10/infra/router/gin"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Init(router gin.IRouter, db *sqlx.DB, repo Repository) {

	svc := newService(repo)
	authHandler := newHandler(svc)

	ginMiddleware := infragin.NewMiddleware()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authHandler.register)
		authGroup.POST("/login", authHandler.login)
		authGroup.GET("/photo/getAll", authHandler.findPhotoAll)
		authGroup.GET("/comment/getAll", authHandler.findCommentAll)
		authGroup.GET("/media/getAll", authHandler.findMediaAll)

		authGroup.Use(ginMiddleware.Authorization())

		authGroup.POST("/photo/CreatePhoto", authHandler.createPhoto)
		authGroup.GET("/photo/getPhotoByUserId", ginMiddleware.CheckRole([]string{"user", "admin"}), authHandler.findPhotoByUserId)
		authGroup.PUT("/photo/updatePhotoByUserId", authHandler.updatePhotoByUserId)
		authGroup.DELETE("/photo/deletePhotoByUserId", authHandler.deletePhotoByUserId)

		authGroup.POST("/comment/createComment", authHandler.createComment)
		authGroup.GET("/comment/getCommentByUserId", ginMiddleware.CheckRole([]string{"user", "admin"}), authHandler.findCommentByUserId)
		authGroup.PUT("/comment/updateCommentByUserId", authHandler.updateCommentByUserId)
		authGroup.DELETE("/comment/deleteCommentByUserId", authHandler.deleteCommentByUserId)

		authGroup.POST("/media/createMedia", authHandler.createMedia)
		authGroup.GET("/media/getMediaByUserId", ginMiddleware.CheckRole([]string{"user", "admin"}), authHandler.findMediaByUserId)
		authGroup.PUT("/media/updateMediaByUserId", authHandler.updateMediaByUserId)
		authGroup.DELETE("/media/deleteMediatByUserId", authHandler.deleteMediaByUserId)

	}
}
