package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"todoList/internal/common/configs"
	"todoList/internal/handlers"
	"todoList/internal/repository"
	"todoList/internal/repository/connection"
	"todoList/internal/service"
)

func Run() error {
	configs.LoadConfig()
	connect, err := connection.GetPostgresConnection(configs.AppConfig.DBConfig)
	if err != nil {
		return err
	}

	repo := repository.New(connect)
	services := service.NewService(repo)
	handler := handlers.New(services)

	app := fiber.New()
	handler.Init(app)

	err = app.Listen(fmt.Sprintf(":%s", configs.AppConfig.Port))

	return err
}
