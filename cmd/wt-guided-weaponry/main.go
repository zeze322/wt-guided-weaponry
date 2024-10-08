package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/zeze322/wt-guided-weaponry/internal/api"
	"github.com/zeze322/wt-guided-weaponry/internal/db/mongodb"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load env file")
	}

	var (
		port            = os.Getenv("PORT")
		mongoURI        = os.Getenv("MONGO_URI")
		mongoDatabase   = os.Getenv("MONGODB_DATABASE")
		mongoCollection = os.Getenv("MONGODB_COLLECTION")
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongodb.New(ctx, mongoURI, mongoDatabase, mongoCollection)
	if err != nil {
		log.Fatal(err)
	}

	if err := mongoClient.CreateIndex(ctx); err != nil {
		log.Fatal(err)
	}

	defer mongoClient.Close(ctx)

	server := api.NewServer(port, mongoClient)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
