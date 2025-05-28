package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-postgres-test/infrastructure"
	"go-postgres-test/internal/delivery/http"
	"go-postgres-test/internal/repository"
	"go-postgres-test/internal/usecase"
)

func main() {
	db := infrastructure.ConnectDB()
	defer db.Close(nil)

	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := http.NewUserHandler(userUC)

	r := gin.Default()
	r.Use(cors.Default())

	userHandler.RegisterRoutes(r)

	fmt.Println("Running server")
	r.Run(":8080")
}
