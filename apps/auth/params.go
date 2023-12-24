package auth

import (
	"errors"
	"regexp"
)

var (
	ErrorPayloadUsernameRequired  = errors.New("Username is required")
	ErrorPayloadEmailRequired     = errors.New("email is required")
	ErrorPayloadEmailInvalid      = errors.New("email is invalid")
	ErrorPayloadPasswordRequired  = errors.New("password is required")
	ErrorPayloadPasswordInvalid   = errors.New("password is invalid")
	ErrorPayloadAgeRequired       = errors.New("age is required")
	ErrorPayloadAgeInvalid        = errors.New("age is invalid > 8")
	ErrorPayloadTitleRequired     = errors.New("title is required")
	ErrorPayloadPhotoRequired     = errors.New("photo_url is required")
	ErrorPayloadMessageRequired   = errors.New("message is required")
	ErrorPayloadMediaNameRequired = errors.New("social media name is required")
	ErrorPayloadMediaUrlRequired  = errors.New("social media url is required")
)

type RegisterPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func (r RegisterPayload) ToAuthEntity() AuthEntity {
	return AuthEntity{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		Age:      r.Age,
	}
}

func (r RegisterPayload) Validate() error {

	if r.Username == "" {
		return ErrorPayloadUsernameRequired
	}

	if r.Email == "" {
		return ErrorPayloadEmailRequired
	}

	if len(r.Email) < 6 {
		return ErrorPayloadEmailInvalid
	}

	if r.Password == "" {
		return ErrorPayloadPasswordRequired
	}

	if len(r.Password) < 6 {
		return ErrorPayloadPasswordInvalid
	}

	if r.Age == 0 {
		return ErrorPayloadAgeRequired
	}

	if r.Age < 8 {
		return ErrorPayloadAgeInvalid
	}

	return nil
}

type LoginRequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (r LoginRequestPayload) Validate() error {
	if r.Username == "" {
		return ErrorPayloadUsernameRequired
	}

	if r.Password == "" {
		return ErrorPayloadPasswordRequired
	}

	if len(r.Password) < 6 {
		return ErrorPayloadPasswordInvalid
	}
	return nil
}

func isValidEmailSetup(email string) bool {

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return false
	}
	return match
}

type CreatePhotoPayload struct {
	Title    string `json:"title" db:"title"`
	Caption  string `json:"caption" db:"caption"`
	PhotoUrl string `json:"photoUrl" db:"photo_url"`
	UserId   int    `json:"userId" db:"user_id"`
}

func (r CreatePhotoPayload) ValidatePhoto() error {

	if r.Title == "" {
		return ErrorPayloadTitleRequired
	}

	if r.PhotoUrl == "" {
		return ErrorPayloadPhotoRequired
	}

	return nil
}

type CreateCommentPayload struct {
	UserId  int    `json:"userId" db:"user_id"`
	PhotoId string `json:"photoId" db:"photo_id"`
	Message string `json:"message" db:"message"`
}

func (r CreateCommentPayload) ValidateComment() error {

	if r.Message == "" {
		return ErrorPayloadMessageRequired
	}

	return nil
}

type CreateMediaPayload struct {
	Name           string `json:"name" db:"name"`
	UserId         int    `json:"userId" db:"user_id"`
	SocialMediaUrl string `json:"socialMediaUrl" db:"social_media_url"`
}

func (r CreateMediaPayload) ValidateMedia() error {

	if r.Name == "" {
		return ErrorPayloadMediaNameRequired
	}

	if r.SocialMediaUrl == "" {
		return ErrorPayloadMediaUrlRequired
	}

	return nil
}
