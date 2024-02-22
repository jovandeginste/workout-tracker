package main

import (
	"embed"
	"time"

	"github.com/jovandeginste/workout-tracker/pkg/app"
	"github.com/labstack/echo/v4"
)

var (
	//go:embed assets/*
	assets   embed.FS
	AssetsFS = echo.MustSubFS(assets, "assets")

	//go:embed views/*
	views   embed.FS
	ViewsFS = echo.MustSubFS(views, "views")

	//go:embed translations/*
	translations   embed.FS
	TranslationsFS = echo.MustSubFS(translations, "translations")

	gitRef     = "0.0.0-dev"
	gitRefName = "local"
	gitRefType = "local"
	gitCommit  = "local"
	buildTime  = time.Now().Format(time.RFC3339)
)

func main() {
	a := app.NewApp(app.Version{
		BuildTime: buildTime,
		Ref:       gitRef,
		RefName:   gitRefName,
		RefType:   gitRefType,
		Sha:       gitCommit,
	})
	a.Assets = AssetsFS
	a.Views = ViewsFS
	a.Translations = TranslationsFS

	if err := a.Configure(); err != nil {
		panic(err)
	}

	if err := a.Serve(); err != nil {
		panic(err)
	}
}
