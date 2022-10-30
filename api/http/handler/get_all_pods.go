package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rra696/kuber-barn/internal/services/pod"
)

func GetAllPodsHandler(c echo.Context) error {
	pods, err := pod.ListRunningPods()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, pods)
}
