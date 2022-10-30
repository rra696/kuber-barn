package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rra696/kuber-barn/internal/services/pod"
)

func DeletePodHandler(c echo.Context) error {
	_, err := pod.KillPod(c.Param("id"))
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
