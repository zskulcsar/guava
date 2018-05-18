package handlers

import (
  "os"

  "net/http"

  // "image"
  // "image/png"

  "github.com/labstack/echo"
)

func AvatarHandler(ctx echo.Context) error {
  //email := ctx.Param("email")

  ir, err := os.Open("static/image.png")
  if err != nil {
    return err
  }
  defer ir.Close()

  return ctx.Stream(http.StatusOK, "image/png", ir)
}