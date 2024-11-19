package server

import (
	"library_app/internal/routes"

	"github.com/gin-gonic/gin"
)

type application struct {
	engine *gin.Engine
}

func (app *application) Run() {
	if err := routes.SetupRouter(app.engine); err != nil {
		panic("Application error")
	}
}

func NewServer() *application {
	router := gin.Default()

	return &application{
		engine: router,
	}

}
