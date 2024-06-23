package main

import (
	library "github.com/automacon-gromoff/FinalProject"
	"github.com/automacon-gromoff/FinalProject/pkg/handler"
	"github.com/automacon-gromoff/FinalProject/pkg/repository"
	"github.com/automacon-gromoff/FinalProject/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	// устанавливаем логирование в формате JSON
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// считываем файл конфигурации
	if err := initConfig(); err != nil {
		logrus.Fatalf("ошибка инициализации конфигурации %s", err.Error())
	}

	// получаем переменные окружения
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("ошибка получения переменной окружения %s", err.Error())
	}

	// подключаемся к базе данных
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: viper.GetString("db.database"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("ошибка инициализации базы данных: %s", err.Error())
	}

	// запускаем http-сервер
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(library.Server)
	if err := srv.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("ошибка запуска http-сервера: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
