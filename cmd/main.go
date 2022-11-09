package main

import (
	"github.com/Krukiscookie/intern_task/internal/handler"
	repository2 "github.com/Krukiscookie/intern_task/internal/repository"
	"github.com/Krukiscookie/intern_task/internal/service"
	"github.com/Krukiscookie/intern_task/pkg/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

//@title User balance API
//@version 1.0.0
//@description Microservice for working with user balance

//@host localhost:8000
//@BasePath /

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := InitConf(); err != nil {
		logrus.Fatalf("error initialazing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env: %s", err.Error())
	}

	db, err := repository2.NewDB(repository2.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository2.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(models.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured  while running http server: %s", err.Error())
	}
}

func InitConf() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
