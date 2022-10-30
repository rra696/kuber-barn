package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rra696/kuber-barn/internal/services/pod"
)

type podCreatingDTO struct {
	ImageRegistry string `json:"image registry"`
	Name          string `json:"name"`
}

func CreatePodHandler(c echo.Context) error {
	req := new(podCreatingDTO)

	err := c.Bind(req)
	if err != nil {
		return err
	}

	id, err := pod.NewPodAndRun(req.ImageRegistry, req.Name)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, podCreatingDTO{
		ImageRegistry: req.ImageRegistry,
		Name:          id,
	})
}
