package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rra696/kuber-barn/internal/services/pod"
)

func LogPodHandler(c echo.Context) error {
	logs, err := pod.LogPod(c.Param("id"))
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, logs)
}
