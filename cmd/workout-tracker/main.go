package main

import (
	appassets "github.com/jovandeginste/workout-tracker/v2/assets"
	"github.com/jovandeginste/workout-tracker/v2/pkg/app"
	"github.com/jovandeginste/workout-tracker/v2/pkg/version"
	apptranslations "github.com/jovandeginste/workout-tracker/v2/translations"
)

var (
	gitRef     = "0.0.0-dev"
	gitRefName = "local"
	gitRefType = "local"
	gitCommit  = "local"
	buildTime  = "now"
)

func main() {
	a := app.NewApp(version.Version{
		BuildTime: buildTime,
		Ref:       gitRef,
		RefName:   gitRefName,
		RefType:   gitRefType,
		Sha:       gitCommit,
	})
	a.Assets = appassets.FS()
	a.Translations = apptranslations.FS()

	if err := a.Configure(); err != nil {
		panic(err)
	}

	if err := a.Serve(); err != nil {
		panic(err)
	}
}
