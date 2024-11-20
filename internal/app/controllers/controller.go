package controllers

import "github.com/mayuka-c/leaderboard-system-go/internal/app/handlers"

type Application struct {
	handler handlers.IServiceHandler
}

func NewApplication() *Application {
	return &Application{
		handler: handlers.NewServiceHandler(),
	}
}
