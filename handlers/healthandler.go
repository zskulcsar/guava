package handlers

import (
  "net/http"

  "github.com/labstack/echo"
)

func HealthHandler(ctx echo.Context) error {
  return ctx.String(http.StatusOK, http.StatusText(http.StatusOK))
}