package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bookstore/docs"
	bookRepo "bookstore/domain/book/repository"
	orderRepo "bookstore/domain/order/repository"
	userRepo "bookstore/domain/user/repository"
	httpSwagger "github.com/swaggo/http-swagger"

	"bookstore/config"
	"bookstore/domain/book"
	"bookstore/domain/order"
	"bookstore/domain/user"
	"bookstore/handler"
	"bookstore/infra/database"
	"bookstore/infra/router"
)

// @title		Bookstore API
// @description	Bookstore API Service
// @host		localhost:9001
// @BasePath	/
func main() {
	ctx := context.Background()
	cfg := config.NewConfig(ctx)

	db, err := database.NewPostgreSQL(cfg.PostgresConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %s\n", err)
	}

	userSvc := user.NewService(userRepo.NewRepository(db))
	bookSvc := book.NewService(bookRepo.NewRepository(db))
	orderSvc := order.NewService(orderRepo.NewRepository(db))

	h := handler.NewHandler(userSvc, bookSvc, orderSvc)
	docs.SwaggerInfo.Host = cfg.Host
	r := router.NewRouter(h)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", cfg.BaseURL)),
	))

	Serve(cfg, r)
}

func Serve(cfg *config.Config, handler http.Handler) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: handler,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("cant start the server: %s\n", err)
		}
	}()
	log.Printf("listening to port :%d", cfg.Port)
	log.Printf("swagger can be accessed at: %s/swagger/index.html", cfg.BaseURL)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown:", err)
	}
	log.Println("server exiting")
}
