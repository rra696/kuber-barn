package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rra696/kuber-barn/api/http/handler"
)

const PORT = "6050"

func main() {
	startContainerd()

	e := echo.New()

	initMiddlewares(e)
	initRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", PORT)))
}

func startContainerd() {
	cmd := exec.Command("containerd")
	cmd.Stdout = os.Stdout

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("containerd run on %d", cmd.Process.Pid)
}

func initMiddlewares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
		},
	))

	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		ctx.Logger().Error(err)

		e.DefaultHTTPErrorHandler(err, ctx)
	}
}

func initRoutes(e *echo.Echo) {
	e.POST("/pods", handler.CreatePodHandler)
	e.GET("/pods/:id/log", handler.LogPodHandler)
	e.GET("/pods", handler.GetAllPodsHandler)
	e.DELETE("pods/:id", handler.DeletePodHandler)
}
