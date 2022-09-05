package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/jungmu/go-web/ent"
)

var db *ent.Client

func init() {
	const rootDir = "go-web"
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	lastIndex := strings.LastIndex(dir, rootDir)
	lastIndex += len(rootDir)
	rootPath := dir[:lastIndex]
	err = os.Chdir(rootPath)
	if err != nil {
		panic(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbname, password))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	db = client
}

func Client() *ent.Client {
	return db
}
