package main

import (
	"log"
	"sesi-10/apps/auth"
	database "sesi-10/external/db"
	infraRepository "sesi-10/infra/repository"
	"sesi-10/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	err := config.Load("cmd/jwt/config.yaml")
	if err != nil {
		log.Println("file config not found")
	} else {
		log.Println("Config loaded successfully:", config.Cfg)
	}

	db, err := database.ConnectDatabase(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := infraRepository.NewRepository(db)

	auth.Init(router, db, repo)

	router.Run(":4444")
}
