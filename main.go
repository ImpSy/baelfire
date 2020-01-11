package main

import (
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gopkg.in/alecthomas/kingpin.v2"
)

var version string

func main() {
	var (
		dbDir  = kingpin.Flag("db.path", "Base path for db storage.").Default("data/").String()
		dbName = kingpin.Flag("db.name", "File name for db file.").Default("baelfire.db").String()

		listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for the web interface and API.").Default(":1323").String()
	)

	kingpin.Version(version)
	kingpin.CommandLine.GetFlag("help").Short('h')
	kingpin.Parse()

	db := createDB(*dbDir, *dbName)
	defer db.close()

	handler := handler{
		db: db,
	}

	// Echo instance
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	api := e.Group("/api/v1")
	api.GET("/version", getVersion)
	api.GET("/targets", handler.listTargets)
	api.POST("/targets", handler.createTarget)
	api.GET("/targets/:name", handler.getTarget)
	api.DELETE("/targets/:name", handler.deleteTarget)
	api.GET("/targets/:name/version", handler.getTargetVersion)

	// Start server
	e.Logger.Fatal(e.Start(*listenAddress))
}

func getVersion(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"version":   version,
		"goVersion": runtime.Version(),
	})
}
