package config

import (
	"context"
	"ecom/db"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"log"
	"os"
)

type Config struct {
	AppPort           string
	UserCollection    *mongo.Collection
	ProductCollection *mongo.Collection
}

func Configurations() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("/***************************** No .env file found *****************************/")
	}

	ctx := context.Background()
	mongoURI := getEnv("MONGO_URI", "mongodb://localhost:27017")
	dbName := getEnv("DATABASE_NAME", "staffy")
	userColl := getEnv("USER_COLLECTION", "users")
	productColl := getEnv("PRODUCT_COLLECTION", "products")
	appPort := getEnv("APP_PORT", "8080")

	client, err := db.ConnectMongo(mongoURI)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	// Ping the database
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("MongoDB ping error:", err)
	}

	fmt.Println("/***************************** Connected to MongoDB! *****************************/")

	return &Config{
		AppPort:           appPort,
		UserCollection:    client.Database(dbName).Collection(userColl),
		ProductCollection: client.Database(dbName).Collection(productColl),
	}
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
