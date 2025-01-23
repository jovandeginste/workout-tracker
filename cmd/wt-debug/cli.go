package main

import (
	appassets "github.com/jovandeginste/workout-tracker/v2/assets"
	"github.com/jovandeginste/workout-tracker/v2/pkg/app"
	"github.com/jovandeginste/workout-tracker/v2/pkg/version"
	apptranslations "github.com/jovandeginste/workout-tracker/v2/translations"
	"gorm.io/gorm"
)

type cli struct {
	app *app.App
}

func newCLI() (*cli, error) {
	a := app.NewApp(version.Version{
		BuildTime: buildTime,
		Ref:       gitRef,
		RefName:   gitRefName,
		RefType:   gitRefType,
		Sha:       gitCommit,
	})
	a.Assets = appassets.FS()
	a.Translations = apptranslations.FS()

	if err := a.ReadConfiguration(); err != nil {
		return nil, err
	}

	a.ConfigureLogger()

	if err := a.ConfigureDatabase(); err != nil {
		return nil, err
	}

	c := &cli{
		app: a,
	}

	return c, nil
}

func (c *cli) getDatabase() *gorm.DB {
	return c.app.DB()
}
