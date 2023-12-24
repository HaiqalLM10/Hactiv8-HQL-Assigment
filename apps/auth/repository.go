package auth

import (
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

var (
	ErrorRepositoryNotFound = errors.New("repository not found")
)

type repository struct {
	db *sqlx.DB
}

// CreateAuth implements Repository.
func (r repository) CreateAuth(req AuthEntity) error {
	query := `
		INSERT INTO "user"(username, email, password, age)
		VALUES(:username, :email, :password, :age)
	`
	_, err := r.db.NamedExec(query, req)
	return err
}

// IsEmailAlreadyExists implements Repository.
func (r repository) IsEmailAlreadyExists(email string) (AuthEntity, error) {

	log.Println("Start Query")
	query := `
		SELECT
			id, username, email, password
		FROM "user"
		WHERE username=$1
	`

	log.Println("Finish Query")

	var authEntity = AuthEntity{}
	err := r.db.Get(&authEntity, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return authEntity, ErrorRepositoryNotFound
		}
		return authEntity, err
	}
	return authEntity, err
}

func (r repository) GetAuthById(id int) (authEntity AuthEntity, err error) {
	query := `
		SELECT
			id, username, password, age
		FROM auth
		WHERE id=$1
	`
	err = r.db.Get(&authEntity, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return authEntity, ErrorRepositoryNotFound
		}
		return authEntity, err
	}
	return
}

func newRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreatePhotoRep(req CreatePhotoPayload) error {

	log.Println("start repo.go")
	query := `
		INSERT INTO "photo"(title, caption, photo_url, user_id)
		VALUES(:title, :caption, :photo_url, :user_id)
	`
	_, err := r.db.NamedExec(query, req)
	return err
}

func (r repository) GetPhotoAll() (photoEntity PhotoEntity, err error) {
	query := `
		SELECT
			title, caption, photo_url, user_id
		FROM photo
	`
	err = r.db.Get(&photoEntity, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return photoEntity, ErrorRepositoryNotFound
		}
		return photoEntity, err
	}
	return
}

func (r repository) GetPhotoByUserId(id int) (photoEntity PhotoEntity, err error) {
	query := `
		SELECT
			title, caption, photo_url, user_id
		FROM photo
		WHERE user_id = $1
	`
	err = r.db.Get(&photoEntity, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return photoEntity, ErrorRepositoryNotFound
		}
		return photoEntity, err
	}
	return
}

func (r repository) UpdatePhotoByUserId(req CreatePhotoPayload) error {

	query := `
		UPDATE "photo" SET
		caption = :caption, 
		photo_url = :photo_url, 
		updated_at = current_timestamp
		WHERE user_id = :user_id
		AND title = :title
	`

	_, err := r.db.NamedExec(query, req)
	return err
}

func (r repository) DeletePhotoByUserId(req CreatePhotoPayload) error {

	query := `
		DELETE from "photo" 
		WHERE user_id = :user_id
		AND title = :title
	`
	_, err := r.db.NamedExec(query, req)
	return err
}

func (r repository) CreateCommentRep(req CreateCommentPayload) error {

	query := `
		INSERT INTO "comment"(user_id, photo_id, message)
		VALUES(:user_id, :photo_id, :message)
	`
	_, err := r.db.NamedExec(query, req)
	return err
}

func (r repository) GetCommentAll() (commentEntity CommentEntity, err error) {
	query := `
		SELECT
			user_id, photo_id, message
		FROM comment
	`
	err = r.db.Get(&commentEntity, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return commentEntity, ErrorRepositoryNotFound
		}
		return commentEntity, err
	}
	return
}

func (r repository) GetCommentByUserId(id int) (commentEntity CommentEntity, err error) {
	query := `
		SELECT
			user_id, photo_id, message
		FROM comment
		WHERE user_id = $1
	`
	err = r.db.Get(&commentEntity, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return commentEntity, ErrorRepositoryNotFound
		}
		return commentEntity, err
	}
	return
}

func (r repository) UpdateCommentByUserId(req CreateCommentPayload) error {

	query := `
		UPDATE "comment" SET
		message = :message, 
		updated_at = current_timestamp
		WHERE user_id = :user_id
		AND photo_id = :photo_id
	`

	_, err := r.db.NamedExec(query, req)
	return err
}

func (r repository) DeleteCommentByUserId(req CreateCommentPayload) error {

	query := `
		DELETE from "comment" 
		WHERE user_id = :user_id
		AND photo_id = :photo_id
	`
	_, err := r.db.NamedExec(query, req)
	return err
}

func (r repository) CreateMediaRep(req CreateMediaPayload) error {

	query := `
		INSERT INTO "socialmedia"(name, social_media_url, user_id)
		VALUES(:name, :social_media_url, :user_id)
	`
	_, err := r.db.NamedExec(query, req)
	return err
}

func (r repository) GetMediaAll() (mediaEntity MediaEntity, err error) {
	query := `
		SELECT
			user_id, name, social_media_url
		FROM socialmedia
	`
	err = r.db.Get(&mediaEntity, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return mediaEntity, ErrorRepositoryNotFound
		}
		return mediaEntity, err
	}
	return
}

func (r repository) GetMediaByUserId(id int) (mediaEntity MediaEntity, err error) {
	query := `
		SELECT
			user_id, name, social_media_url
		FROM socialmedia
		WHERE user_id = $1
	`
	err = r.db.Get(&mediaEntity, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return mediaEntity, ErrorRepositoryNotFound
		}
		return mediaEntity, err
	}
	return
}

func (r repository) UpdateMediaByUserId(req CreateMediaPayload) error {

	query := `
		UPDATE "socialmedia" SET
		name = :name, 
		social_media_url :social_media_url,
		updated_at = current_timestamp
		WHERE user_id = :user_id
	`

	_, err := r.db.NamedExec(query, req)
	return err
}

func (r repository) DeleteMediaByUserId(req CreateMediaPayload) error {

	query := `
		DELETE from "socialmedia" 
		WHERE user_id = :user_id
	`
	_, err := r.db.NamedExec(query, req)
	return err
}
