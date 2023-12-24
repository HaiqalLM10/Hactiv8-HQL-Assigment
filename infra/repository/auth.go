package infraRepository

import (
	"database/sql"
	"sesi-10/apps/auth"
)

// CreateAuth implements Repository.
func (r repository) CreateAuth(req auth.AuthEntity) error {
	query := `
		INSERT INTO auth(email, password)
		VALUES(:email, :password)
	`
	_, err := r.db.NamedExec(query, req)
	return err
}

// IsEmailAlreadyExists implements Repository.
func (r repository) IsEmailAlreadyExists(email string) (auth.AuthEntity, error) {
	query := `
		SELECT
			id, email, password
		FROM auth
		WHERE email=$1
	`

	var authEntity = auth.AuthEntity{}
	err := r.db.Get(&authEntity, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return authEntity, ErrorRepositoryNotFound
		}
		return authEntity, err
	}
	return authEntity, err
}

func (r repository) GetAuthById(id int) (authEntity auth.AuthEntity, err error) {
	query := `
		SELECT
			id, email, password
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

func (r repository) CreatePhotoRep(req auth.CreatePhotoPayload) error {

	query := `
		INSERT INTO "photo" (title, caption, photo_url, user_id)
		VALUES (:title, :caption, :photo_url, :user_id)
	`
	_, err := r.db.NamedExec(query, req)
	return err
}

func (r repository) GetPhotoAll() (photoEntity auth.PhotoEntity, err error) {
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

func (r repository) GetPhotoByUserId(id int) (photoEntity auth.PhotoEntity, err error) {
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

func (r repository) UpdatePhotoByUserId(req auth.CreatePhotoPayload) error {

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

func (r repository) DeletePhotoByUserId(req auth.CreatePhotoPayload) error {

	query := `
		DELETE from "photo" 
		WHERE user_id = :user_id
		AND title = :title
	`
	_, err := r.db.NamedExec(query, req)
	return err
}

func (r repository) CreateCommentRep(req auth.CreateCommentPayload) error {

	query := `
		INSERT INTO "comment"(user_id, photo_id, message)
		VALUES(:user_id, :photo_id, :message)
	`
	_, err := r.db.NamedExec(query, req)
	return err
}

func (r repository) GetCommentAll() (commentEntity auth.CommentEntity, err error) {
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

func (r repository) GetCommentByUserId(id int) (commentEntity auth.CommentEntity, err error) {
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

func (r repository) UpdateCommentByUserId(req auth.CreateCommentPayload) error {

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

func (r repository) DeleteCommentByUserId(req auth.CreateCommentPayload) error {

	query := `
		DELETE from "comment" 
		WHERE user_id = :user_id
		AND photo_id = :photo_id
	`
	_, err := r.db.NamedExec(query, req)
	return err
}

func (r repository) CreateMediaRep(req auth.CreateMediaPayload) error {

	query := `
		INSERT INTO "socialmedia"(name, social_media_url, user_id)
		VALUES(:name, :social_media_url, :user_id)
	`
	_, err := r.db.NamedExec(query, req)
	return err
}

func (r repository) GetMediaAll() (mediaEntity auth.MediaEntity, err error) {
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

func (r repository) GetMediaByUserId(id int) (mediaEntity auth.MediaEntity, err error) {
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

func (r repository) UpdateMediaByUserId(req auth.CreateMediaPayload) error {

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

func (r repository) DeleteMediaByUserId(req auth.CreateMediaPayload) error {

	query := `
		DELETE from "socialmedia" 
		WHERE user_id = :user_id
	`
	_, err := r.db.NamedExec(query, req)
	return err
}
