package main

import (
	"context"
	"fmt"
	"learning-hub/internal/category"
	"learning-hub/internal/material"
	"learning-hub/internal/reviews"
	"learning-hub/internal/user"
	"net/http"
)

func main() {

	ctx := context.Background()

	conn, err := ConnectDB(ctx)
	if err != nil {
		fmt.Println("DB connected error!", err)
		return
	}
	defer conn.Close(ctx)

	repository := user.NewRepository(conn)

	service := user.NewService(repository)

	handler := user.NewHandler(service)

	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)

	categoryRepository := category.NewRepository(conn)

	categoryService := category.NewService(categoryRepository)

	categoryHandler := category.NewHandler(categoryService)

	http.HandleFunc("/categories", categoryHandler.GetCategories)
	http.HandleFunc("/categories/create", categoryHandler.CreateCategory)

	materialRepository := material.NewRepository(conn)

	materialService := material.NewService(materialRepository)

	materialHandler := material.NewHandler(materialService)

	http.HandleFunc("/materials", materialHandler.GetMaterials)
	http.HandleFunc("/materials/create", materialHandler.CreateMaterial)

	reviewRepository := reviews.NewRepository(conn)

	reviewService := reviews.NewService(reviewRepository)

	reviewHandler := reviews.NewHandler(reviewService)

	http.HandleFunc("/reviews", reviewHandler.GetReviews)
	http.HandleFunc("/reviews/create", reviewHandler.CreateReview)
	http.HandleFunc("/reviews/delete", reviewHandler.DeleteReview)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error", err)
	}

	// http.Handle("/profile", handler.JWTMiddleware(http.HandlerFunc(handler.Profile)))

}
