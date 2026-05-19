package apptranslations

import (
	"bytes"
	"embed"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

//go:embed *
var embedded embed.FS

type rewriteFS struct {
	fs.FS
}

func (r rewriteFS) ReadFile(name string) ([]byte, error) {
	data, err := fs.ReadFile(r.FS, name)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(name)
	if ext != ".yaml" && ext != ".yml" {
		return data, nil
	}

	base := filepath.Base(name)
	lang := strings.TrimSuffix(base, ext)
	lang = strings.ReplaceAll(lang, "_", "-")

	lines := strings.Split(string(data), "\n")
	var buf bytes.Buffer
	buf.WriteString(lang + ":\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "---" {
			continue
		}
		if line == "" {
			buf.WriteString("\n")
		} else {
			buf.WriteString("  " + line + "\n")
		}
	}
	return buf.Bytes(), nil
}

func FS() fs.FS {
	subFS := echo.MustSubFS(embedded, "")
	return rewriteFS{subFS}
}
