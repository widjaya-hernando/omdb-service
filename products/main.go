package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/widjaya-hernando/prod-trans-user-service/lib/database_transaction"
	"github.com/widjaya-hernando/prod-trans-user-service/products/repository"
	"github.com/widjaya-hernando/prod-trans-user-service/products/server"

	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	loadEnv()
}

func main() {
	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()

	db := setupDatabase()
	var tM database_transaction.Client
	repo := repository.NewRepo(db)

	srv, err := server.NewServer(ctx, db, repo, tM)
	if err != nil {
		log.Fatalf("NewServer failed with error: %s", err)
	}

	srv.Run()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	signal := <-sigChan
	log.Printf("shutting down products server with signal: %s", signal)
}

func loadEnv() {
	err := gotenv.Load("../.env")

	if err != nil {
		log.Println("failed to load from .env")
	}
}

func setupDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		// "localhost",
		// "postgres",
		// "postgres",
		// "prod-trans-user-service",
		// "5432",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
	)

	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
		//Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	return db
}
