package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) register(ctx *gin.Context) {
	var req = RegisterPayload{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	err := h.svc.Register(req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorEmailAlreadyExists:
			statusCode = http.StatusBadRequest
		}
		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (h handler) login(ctx *gin.Context) {
	var req = LoginRequestPayload{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	token, err := h.svc.Login(req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"payload": map[string]interface{}{
			"access_token": token,
		},
	})

}

func (h handler) createPhoto(ctx *gin.Context) {

	authId := ctx.GetInt("auth_id")
	if authId == 0 {
		log.Println("auth_id not provided")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	var req = CreatePhotoPayload{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := req.ValidatePhoto(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	err := h.svc.CreatePhotoService(authId, req)

	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success Create Photo",
	})
}

func (h handler) findPhotoAll(ctx *gin.Context) {

	authEntity, err := h.svc.getPhotoAll()
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"payload": authEntity.ParseToProfileResponsePhoto(),
	})
}

func (h handler) findPhotoByUserId(ctx *gin.Context) {

	authId := ctx.GetInt("auth_id")
	log.Println("authId: ", authId)
	if authId == 0 {
		log.Println("auth_id not provided")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	authEntity, err := h.svc.getPhotoAll()
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"payload": authEntity.ParseToProfileResponsePhoto(),
	})
}

func (h handler) updatePhotoByUserId(ctx *gin.Context) {

	authId := ctx.GetInt("auth_id")
	if authId == 0 {
		log.Println("auth_id not provided")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	var req = CreatePhotoPayload{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := req.ValidatePhoto(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	err := h.svc.updatePhotoByUserId(authId, req)

	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success Update Photo",
	})
}

func (h handler) deletePhotoByUserId(ctx *gin.Context) {

	authId := ctx.GetInt("auth_id")
	if authId == 0 {
		log.Println("auth_id not provided")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	var req = CreatePhotoPayload{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if req.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Title is Mandatory",
		})
		return
	}

	err := h.svc.deletePhotoByUserId(authId, req)

	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success Delete Photo",
	})
}

func (h handler) createComment(ctx *gin.Context) {

	authId := ctx.GetInt("auth_id")
	if authId == 0 {
		log.Println("auth_id not provided")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	var req = CreateCommentPayload{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := req.ValidateComment(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	err := h.svc.CreateCommentService(authId, req)

	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success Create Comment",
	})
}

func (h handler) findCommentAll(ctx *gin.Context) {

	authEntity, err := h.svc.getCommentAll()
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"payload": authEntity.ParseToProfileResponseComment(),
	})
}

func (h handler) findCommentByUserId(ctx *gin.Context) {

	authId := ctx.GetInt("auth_id")
	if authId == 0 {
		log.Println("auth_id not provided")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	authEntity, err := h.svc.getCommentByUserId(authId)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"payload": authEntity.ParseToProfileResponseComment(),
	})
}

func (h handler) updateCommentByUserId(ctx *gin.Context) {

	authId := ctx.GetInt("auth_id")
	if authId == 0 {
		log.Println("auth_id not provided")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	var req = CreateCommentPayload{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := req.ValidateComment(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	err := h.svc.updateCommentByUserId(authId, req)

	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success Update Comment",
	})
}

func (h handler) deleteCommentByUserId(ctx *gin.Context) {

	authId := ctx.GetInt("auth_id")
	if authId == 0 {
		log.Println("auth_id not provided")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	var req = CreateCommentPayload{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if req.PhotoId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Photo ID is Mandatory",
		})
		return
	}

	err := h.svc.deleteCommentByUserId(authId, req)

	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success Delete Comment",
	})
}

func (h handler) createMedia(ctx *gin.Context) {

	authId := ctx.GetInt("auth_id")
	if authId == 0 {
		log.Println("auth_id not provided")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	var req = CreateMediaPayload{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := req.ValidateMedia(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	err := h.svc.CreateMediaService(authId, req)

	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success Create Media",
	})
}

func (h handler) findMediaAll(ctx *gin.Context) {

	authEntity, err := h.svc.getCommentAll()
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"payload": authEntity.ParseToProfileResponseComment(),
	})
}

func (h handler) findMediaByUserId(ctx *gin.Context) {

	authId := ctx.GetInt("auth_id")
	if authId == 0 {
		log.Println("auth_id not provided")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	authEntity, err := h.svc.getMediaByUserId(authId)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"payload": authEntity.ParseToProfileResponseMedia(),
	})
}

func (h handler) updateMediaByUserId(ctx *gin.Context) {

	authId := ctx.GetInt("auth_id")
	if authId == 0 {
		log.Println("auth_id not provided")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	var req = CreateMediaPayload{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := req.ValidateMedia(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	err := h.svc.updateMediaByUserId(authId, req)

	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success Update Comment",
	})
}

func (h handler) deleteMediaByUserId(ctx *gin.Context) {

	authId := ctx.GetInt("auth_id")
	if authId == 0 {
		log.Println("auth_id not provided")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "unauthorized",
		})
		return
	}

	var req = CreateMediaPayload{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	err := h.svc.deleteMediaByUserId(authId, req)

	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err {
		case ErrorRepositoryNotFound, ErrorPasswordNotMatch:
			statusCode = http.StatusUnauthorized
		}

		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success Delete Photo",
	})
}
