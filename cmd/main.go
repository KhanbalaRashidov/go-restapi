package main

import (
	"context"
	_ "github.com/KhanbalaRashidov/go-restapi"
	"github.com/KhanbalaRashidov/go-restapi/pkg/handler"
	"github.com/KhanbalaRashidov/go-restapi/pkg/repository"
	"github.com/KhanbalaRashidov/go-restapi/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/KhanbalaRashidov/go-restapi/docs"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"

	gorestapi "github.com/KhanbalaRashidov/go-restapi"
)

// @title Go Rest API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("Error initializing database connection: %s", err.Error())
	}

	repos := repository.NeRepository(db)
	services := service.NeService(repos)
	handlers := handler.NewHandler(services)

	srv := new(gorestapi.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error ocurred  while  running  http  server: %s", err.Error())
		}
	}()

	logrus.Printf("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Printf("TodoApp Shutting down ")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
