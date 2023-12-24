package auth

import "log"

type Repository interface {
	IsEmailAlreadyExists(email string) (AuthEntity, error)
	GetAuthById(id int) (AuthEntity, error)
	CreateAuth(req AuthEntity) error
	CreatePhotoRep(req CreatePhotoPayload) error
	GetPhotoAll() (PhotoEntity, error)
	GetPhotoByUserId(id int) (PhotoEntity, error)
	UpdatePhotoByUserId(req CreatePhotoPayload) error
	DeletePhotoByUserId(req CreatePhotoPayload) error
	CreateCommentRep(req CreateCommentPayload) error
	GetCommentAll() (CommentEntity, error)
	GetCommentByUserId(id int) (CommentEntity, error)
	UpdateCommentByUserId(req CreateCommentPayload) error
	DeleteCommentByUserId(req CreateCommentPayload) error
	CreateMediaRep(req CreateMediaPayload) error
	GetMediaAll() (MediaEntity, error)
	GetMediaByUserId(id int) (MediaEntity, error)
	UpdateMediaByUserId(req CreateMediaPayload) error
	DeleteMediaByUserId(req CreateMediaPayload) error
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) Register(req RegisterPayload) (err error) {
	authEntity := NewFromRegister(req)

	if err = authEntity.EncryptPassword(); err != nil {
		return
	}

	if err = s.repo.CreateAuth(authEntity); err != nil {
		return
	}
	return
}

func (s service) Login(req LoginRequestPayload) (token string, err error) {
	authEntity, err := s.repo.IsEmailAlreadyExists(req.Username)
	if err != nil {
		log.Println("error when try to get IsEmailAlreadyExists with detail", err.Error())
		return
	}

	if err = authEntity.ValidatePassword(req.Password); err != nil {
		log.Println("error when try to ValidatePassword with detail", err.Error())
		return
	}

	token, err = authEntity.GenerateToken()
	if err != nil {
		log.Println("error when try to GenerateToken with detail", err.Error())
		return
	}
	return
}

func (s service) CreatePhotoService(id int, req CreatePhotoPayload) (err error) {
	req.UserId = id
	err = s.repo.CreatePhotoRep(req)

	return
}

func (s service) getPhotoAll() (photoEntity PhotoEntity, err error) {
	photoEntity, err = s.repo.GetPhotoAll()
	return
}

func (s service) getPhotoByUserId(id int) (photoEntity PhotoEntity, err error) {
	photoEntity, err = s.repo.GetPhotoByUserId(id)
	return
}

func (s service) updatePhotoByUserId(id int, req CreatePhotoPayload) (err error) {
	req.UserId = id
	err = s.repo.UpdatePhotoByUserId(req)

	return
}

func (s service) deletePhotoByUserId(id int, req CreatePhotoPayload) (err error) {
	req.UserId = id
	err = s.repo.DeletePhotoByUserId(req)

	return
}

func (s service) CreateCommentService(id int, req CreateCommentPayload) (err error) {
	req.UserId = id
	err = s.repo.CreateCommentRep(req)

	return
}

func (s service) getCommentAll() (commentEntity CommentEntity, err error) {
	commentEntity, err = s.repo.GetCommentAll()
	return
}

func (s service) getCommentByUserId(id int) (commentEntity CommentEntity, err error) {
	commentEntity, err = s.repo.GetCommentByUserId(id)
	return
}

func (s service) updateCommentByUserId(id int, req CreateCommentPayload) (err error) {
	req.UserId = id
	err = s.repo.UpdateCommentByUserId(req)

	return
}

func (s service) deleteCommentByUserId(id int, req CreateCommentPayload) (err error) {
	req.UserId = id
	err = s.repo.DeleteCommentByUserId(req)

	return
}

func (s service) CreateMediaService(id int, req CreateMediaPayload) (err error) {
	req.UserId = id
	err = s.repo.CreateMediaRep(req)

	return
}

func (s service) getMediaAll() (mediaEntity MediaEntity, err error) {
	mediaEntity, err = s.repo.GetMediaAll()
	return
}

func (s service) getMediaByUserId(id int) (mediaEntity MediaEntity, err error) {
	mediaEntity, err = s.repo.GetMediaByUserId(id)
	return
}

func (s service) updateMediaByUserId(id int, req CreateMediaPayload) (err error) {
	req.UserId = id
	err = s.repo.UpdateMediaByUserId(req)
	return
}

func (s service) deleteMediaByUserId(id int, req CreateMediaPayload) (err error) {
	req.UserId = id
	err = s.repo.DeleteMediaByUserId(req)

	return
}
