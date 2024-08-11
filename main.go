package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Recipes struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title"`
	Ingredients  string             `json:"ingredients"`
	Servings     string             `json:"servings"`
	Instructions string             `json:"instructions"`
	Password     string             `json:"password"`
	User         string             `json:"user"`
}

var collection *mongo.Collection

func main() {

	godotenv.Load()

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	collection = client.Database("recipe_catalog").Collection("recipe_catalog_log")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("/api/recipes", postRecipes)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func postRecipes(c *fiber.Ctx) error {

	recipe := new(Recipes)

	fmt.Println("Request successfully recieved")

	if err := c.BodyParser(recipe); err != nil {
		return err
	}

	insertResult, err := collection.InsertOne(context.Background(), recipe)
	if err != nil {
		return err
	}

	fmt.Println("Insert Successful")

	recipe.ID = insertResult.InsertedID.(primitive.ObjectID)
	fmt.Println("Returning Data")

	return c.Status(201).JSON(recipe)
}
