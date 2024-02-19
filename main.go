package main

import (
	"embed"

	"github.com/jovandeginste/workouts/pkg/app"
	"github.com/labstack/echo/v4"
)

var (
	//go:embed assets/*
	assets   embed.FS
	AssetsFS = echo.MustSubFS(assets, "assets")

	//go:embed views/*
	views   embed.FS
	ViewsFS = echo.MustSubFS(views, "views")

	version = "0.0.0-dev"
)

func main() {
	a := app.NewApp()
	a.Assets = AssetsFS
	a.Views = ViewsFS
	a.Version = version

	if err := a.Configure(); err != nil {
		panic(err)
	}

	if err := a.Serve(); err != nil {
		panic(err)
	}
}
