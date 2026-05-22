package apptranslations

import (
	"embed"
	"io/fs"

	"github.com/labstack/echo/v4"
)

//go:embed *
var embedded embed.FS

func FS() fs.FS {
	return echo.MustSubFS(embedded, "")
}
