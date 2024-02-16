package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tkrajina/gpxgo/gpx"
)

func uploadedGPXFile(c echo.Context) (*gpx.GPX, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	gpxContent, err := ParseGPX(src)
	if err != nil {
		return nil, err
	}

	return gpxContent, nil
}

func (a *App) addWorkout(c echo.Context) error {
	gpxContent, err := uploadedGPXFile(c)
	if err != nil {
		a.setError(c, err)
		return c.Redirect(http.StatusMovedPermanently, "/workouts/add")
	}

	notes := c.FormValue("notes")

	if err := a.getUser(c).AddWorkout(a.db, notes, gpxContent); err != nil {
		a.setError(c, err)
		return c.Redirect(http.StatusMovedPermanently, "/workouts/add")
	}

	a.setNotice(c, "new workout was added")

	return c.Redirect(http.StatusMovedPermanently, "/workouts")
}
