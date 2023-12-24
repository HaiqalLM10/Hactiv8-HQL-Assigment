package database

import (
	"fmt"
	"log"
	"sesi-10/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDatabase(cfg config.DBConfig) (db *sqlx.DB, err error) {
	log.Println("cfg: ", cfg)
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	if err = Migration(db); err != nil {
		return
	}

	return
}

func Migration(db *sqlx.DB) error {

	query := `
	CREATE TABLE IF NOT EXISTS "user" (
		id SERIAL PRIMARY KEY,
		username varchar(100) NOT NULL,
		email varchar(100) NOT NULL,
		password varchar(255) NOT NULL,
		age int NOT NULL,
		created_at timestamptz DEFAULT NOW(),
		updated_at timestamptz DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS photo (
		id SERIAL PRIMARY KEY,
		title varchar(100) NOT NULL,
		caption varchar(100) NOT NULL,
		photo_url varchar(500) NOT NULL,
		user_id int NOT NULL,
		created_at timestamptz DEFAULT NOW(),
		updated_at timestamptz DEFAULT NOW()
	);

	--ALTER TABLE photo ADD CONSTRAINT photo_FK01
	--FOREIGN KEY(user_id) REFERENCES "user"(id);

	CREATE TABLE IF NOT EXISTS comment (
		id SERIAL PRIMARY KEY,
		user_id int NOT NULL,
		photo_id varchar(100) NOT NULL,
		message varchar(500) NOT NULL,
		created_at timestamptz DEFAULT NOW(),
		updated_at timestamptz DEFAULT NOW()
	);

	--ALTER TABLE comment ADD CONSTRAINT comment_FK01
	--FOREIGN KEY(user_id) REFERENCES "user"(id);

	CREATE TABLE IF NOT EXISTS socialmedia (
		id SERIAL PRIMARY KEY,
		name varchar(100) NOT NULL,
		social_media_url varchar(100) NOT NULL,
		user_id int NOT NULL,
		created_at timestamptz DEFAULT NOW(),
		updated_at timestamptz DEFAULT NOW()
	);

	--ALTER TABLE socialmedia ADD CONSTRAINT socialmedia_FK01
	--FOREIGN KEY(user_id) REFERENCES "user"(id);

	`

	_, err := db.Exec(query)
	return err
}
