package auth

import (
	"errors"
	"sesi-10/internal/config"
	"sesi-10/utility"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorEmailAlreadyExists = errors.New("email already exists")
	ErrorPasswordNotMatch   = errors.New("password not match")
)

type AuthEntity struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Age      int    `db:"age"`
}

func NewFromRegister(req RegisterPayload) AuthEntity {
	return AuthEntity{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Age:      req.Age,
	}
}

func (a *AuthEntity) EncryptPassword() (err error) {

	passEncrypted, err := utility.BcryptHash(a.Password)
	if err != nil {
		return
	}

	a.Password = passEncrypted
	return
}

func (a AuthEntity) ValidatePassword(plain string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plain))
	if err != nil {
		return ErrorPasswordNotMatch
	}
	return
}

func (a AuthEntity) GenerateToken() (token string, err error) {
	return utility.GenerateToken(a.Id, config.Cfg.App.Token)
}

func (a AuthEntity) CheckEmail(email string) (err error) {
	if strings.EqualFold(a.Email, email) {
		return ErrorEmailAlreadyExists
	}
	return
}

func (a AuthEntity) ParseToProfileResponse() ProfileResponse {
	return ProfileResponse{
		Id:    a.Id,
		Email: a.Email,
	}
}

type PhotoEntity struct {
	Id       int    `db:"id"`
	Title    string `db:"title"`
	Caption  string `db:"caption"`
	PhotoUrl string `db:"photo_url"`
	UserId   int    `db:"user_id"`
}

func (a PhotoEntity) ParseToProfileResponsePhoto() PhotoResponse {
	return PhotoResponse{
		Title:    a.Title,
		Caption:  a.Caption,
		PhotoUrl: a.PhotoUrl,
		UserId:   a.UserId,
	}
}

type CommentEntity struct {
	Id      int    `db:"id"`
	UserId  int    `db:"user_id"`
	PhotoId int    `db:"photo_id"`
	Message string `db:"message"`
}

func (a CommentEntity) ParseToProfileResponseComment() CommentEntity {
	return CommentEntity{
		UserId:  a.UserId,
		PhotoId: a.PhotoId,
		Message: a.Message,
	}
}

type MediaEntity struct {
	Id             int    `db:"id"`
	UserId         int    `db:"user_id"`
	Name           string `db:"name"`
	SocialMediaUrl string `db:"social_media_url"`
}

func (a MediaEntity) ParseToProfileResponseMedia() MediaEntity {
	return MediaEntity{
		UserId:         a.UserId,
		Name:           a.Name,
		SocialMediaUrl: a.SocialMediaUrl,
	}
}
