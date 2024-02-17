package app

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
)

func uploadedGPXFile(file *multipart.FileHeader) ([]byte, error) {
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
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["file"]

	msg := ""
	errMsg := ""

	for _, file := range files {
		content, parseErr := uploadedGPXFile(file)
		if parseErr != nil {
			errMsg += parseErr.Error() + "\n"
			continue
		}

		notes := c.FormValue("notes")
		workoutType := c.FormValue("type")

		w, addErr := a.getUser(c).AddWorkout(a.db, workoutType, notes, content)
		if addErr != nil {
			errMsg += addErr.Error() + "\n"
			continue
		}

		msg += "- " + w.Name + "\n"
	}

	if errMsg != "" {
		a.setError(c, errMsg)
	}

	if msg != "" {
		a.setNotice(c, "A new workout was added:\n"+msg)
	}

	return c.Redirect(http.StatusMovedPermanently, "/workouts")
}
