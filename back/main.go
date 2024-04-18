package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/iryoda/price-guru/app/controllers"
	middlewares "github.com/iryoda/price-guru/app/middleware"
	providers_impl "github.com/iryoda/price-guru/app/providers/impl"
	"github.com/iryoda/price-guru/app/repositories/impl"
	"github.com/iryoda/price-guru/app/routes"
	"github.com/iryoda/price-guru/app/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() *mongo.Client {
	mongoURI := os.Getenv("MONGO_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	opts := options.Client().ApplyURI(mongoURI)

	mongo.Connect(ctx, opts)

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func InitDotEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	InitDotEnv()

	client := ConnectToMongo()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	db := client.Database("price-guru")

	// Providers
	tp := providers_impl.JWTProvider{}
	hp := providers_impl.BcryptHashProvider{}

	// Repostiories
	ur := impl.NewUserRepository(db)
	wr := impl.NewMongoWatcherRepository(db)

	// Services
	us := services.UserService{
		UserRepository: &ur,
		TokenProvider:  &tp,
		HashProvider:   &hp,
	}
	ws := services.WatcherService{
		WatcherRepository: &wr,
	}
	as := services.AuthService{
		UserRepository: &ur,
		HashProvider:   &hp,
		TokenProvider:  &tp,
	}

	// Controllers
	uc := controllers.UserController{
		UserService: &us,
	}
	wc := controllers.WatcherController{
		WatcherService: &ws,
	}
	ac := controllers.AuthController{
		AuthService: &as,
		UserService: &us,
	}

	// Routes
	uRouter := routes.UserRouter{
		UserController: &uc,
	}
	wRouter := routes.WatcherRouter{
		WatcherController: &wc,
	}
	aRouter := routes.AuthRouter{
		AuthController: &ac,
	}

	// Middlewares
	authMiddleware := middlewares.AuthMiddleware{
		TokenProvider: &tp,
	}

	// Server
	e := echo.New()

	e.Use(middleware.CORS())

	aRouter.NewAuthRouter(e)

	g := e.Group("/secure")

	g.Use(authMiddleware.WithJWTToken)

	uRouter.NewUserRouter(g)
	wRouter.NewWatcherRouter(g)

	e.Logger.Fatal(e.Start(":1323"))
}
