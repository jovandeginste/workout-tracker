package app

import (
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func uploadedGPXFile(c echo.Context) ([]byte, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	// Read all from r into a bytes slice
	gpxBytes, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}

	return gpxBytes, nil
}

func (a *App) addWorkout(c echo.Context) error {
	content, err := uploadedGPXFile(c)
	if err != nil {
		a.setError(c, err)
		return c.Redirect(http.StatusMovedPermanently, "/workouts/add")
	}

	notes := c.FormValue("notes")

	w, err := a.getUser(c).AddWorkout(a.db, notes, content)
	if err != nil {
		a.setError(c, err)
		return c.Redirect(http.StatusMovedPermanently, "/workouts/add")
	}

	a.setNotice(c, fmt.Sprintf("A new workout was added: %s", w.Name))

	return c.Redirect(http.StatusMovedPermanently, "/workouts")
}
