package main

import (
	"context"
	"fmt"
	"learning-hub/auth"
	"learning-hub/database"
	"learning-hub/internal/category"
	"learning-hub/internal/material"
	"learning-hub/internal/profile"
	"learning-hub/internal/reviews"
	"learning-hub/internal/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	router := gin.Default()

	fmt.Println("Server start!")

	ctx := context.Background()

	conn, err := database.ConnectDB(ctx)
	if err != nil {
		fmt.Println("DB connected error!", err)
		return
	}
	defer conn.Close(ctx)

	//User
	repository := user.NewRepository(conn)
	service := user.NewService(repository, secret)
	handler := user.NewHandler(service)

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	//Category
	categoryRepository := category.NewRepository(conn)
	categoryService := category.NewService(categoryRepository)
	categoryHandler := category.NewHandler(categoryService)

	router.GET("/categories", categoryHandler.GetCategories)
	router.POST("/categories/create", categoryHandler.CreateCategory)

	//Material
	materialRepository := material.NewRepository(conn)
	materialService := material.NewService(materialRepository)
	materialHandler := material.NewHandler(materialService)

	router.GET("/materials", materialHandler.GetMaterials)
	router.POST("/materials/create", materialHandler.CreateMaterial)

	//Review
	reviewRepository := reviews.NewRepository(conn)
	reviewService := reviews.NewService(reviewRepository)
	reviewHandler := reviews.NewHandler(reviewService)

	router.GET("/reviews", reviewHandler.GetReviews)
	router.POST("/reviews/create", reviewHandler.CreateReview)
	router.DELETE("/reviews/:id", reviewHandler.DeleteReview)

	// Profile
	profileRepository := profile.NewRepository(conn)
	profileService := profile.NewService(profileRepository)
	profileHandler := profile.NewHandler(profileService)

	router.PATCH("/profile", auth.Auth(secret), profileHandler.UpdateProfile)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("server error", err)
	}

}
