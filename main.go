package main

import (
	"fmt"
	"os"

	// "encoding/json";
	"log"
	// "net/http";
	// "strings";
	"context"
	"time"

	// "os/signal";
	// "syscall";

	// "github.com/go-chi/cors"
	// "github.com/go-chi/chi";
	// "github.com/go-chi/chi/middleware";
	// "github.com/go-chi/cors";
	"github.com/joho/godotenv"
	"github.com/thedevsaddam/renderer"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "gopkg.in/mgo.v2/bson";
)

var rnd *renderer.Render
var db *mongo.Database

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	rnd = renderer.New()

	srv := os.Getenv("MONGO_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(srv)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}

	// Ping the database to ensure connectivity
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Could not ping MongoDB: %v", err)
	}

	db = client.Database("go-todo") // name of your database
	fmt.Println("Connected to MongoDB")
}

type Todo struct {
	ID   string `json:"id" bson:"_id"`
	Text string `json:"text" bson:"text"`
	Done bool   `json:"done" bson:"done"`
}

type TodoPayload struct {
	Text string `json:"text" bson:"text"`
	Done bool   `json:"done" bson:"done"`
}

func main() {
	fmt.Println("App is starting...")
	// Start server or hold open
	select {} // Keeps the app running for now
}
