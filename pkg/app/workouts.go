package app

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

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

	msg := []string{}
	errMsg := []string{}

	for _, file := range files {
		content, parseErr := uploadedGPXFile(file)
		if parseErr != nil {
			errMsg = append(errMsg, parseErr.Error())
			continue
		}

		notes := c.FormValue("notes")
		workoutType := c.FormValue("type")

		w, addErr := a.getCurrentUser(c).AddWorkout(a.db, workoutType, notes, content)
		if addErr != nil {
			errMsg = append(errMsg, addErr.Error())
			continue
		}

		msg = append(msg, w.Name)
	}

	if len(errMsg) > 0 {
		a.setError(c, fmt.Sprintf("Encountered %d problems while adding workouts: %s", len(errMsg), strings.Join(errMsg, "; ")))
	}

	if len(msg) > 0 {
		a.setNotice(c, fmt.Sprintf("Added %d new workout(s): %s", len(msg), strings.Join(msg, "; ")))
	}

	return c.Redirect(http.StatusFound, a.echo.Reverse("workouts"))
}
