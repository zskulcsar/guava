package cmd

import (
  "time"
  "net/http"
  
  "github.com/spf13/cobra"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"

  "github.com/zsoltk-iw/guava/handlers"
)

var (
  cmdServer = &cobra.Command{
    Use: "server",
    Short: "Start the guava server",
    RunE: runServer,
    SilenceUsage: true,
  }
)

func init() {
  RootCmd.AddCommand(cmdServer)
}

func runServer(cmd *cobra.Command, args []string) error {
  e := echo.New()
  e.Use(middleware.Logger())
  e.Use(middleware.BodyLimit("50M"))

  e.GET("/", redirect)
  e.GET("/avatar/:email", handlers.AvatarHandler)
  e.File("/index", "static/index.html")
  e.File("/robots.txt", "static/robots.txt")

  e.Any("/health", handlers.HealthHandler)

  // Start a custom server
  srv := &http.Server{
    Addr:         gsconf.getServerPort(),
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 20 * time.Second,
  }

  e.Logger.Fatal(e.StartServer(srv))


  return nil
}

// Should be placed somwhere else, but it's so tempting to just chuck it here
func redirect(ctx echo.Context) error {
  return ctx.Redirect(http.StatusPermanentRedirect, "/index")
}