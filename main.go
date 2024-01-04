package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chiboycalix/dns-resolver/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		PassLocalsToViews:     true,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Access-Control-Allow-Credentials",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	initRoutes(app)
	listenAddr := os.Getenv("HTTP_LISTEN_ADDR")
	fmt.Println("app running port " + listenAddr)
	log.Fatal(app.Listen(listenAddr))
}

func initRoutes(app *fiber.App) {
	app.Post("/resolve-dns", handlers.ResolveDNS)
}
