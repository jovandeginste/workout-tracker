package util

import (
	"io/fs"
	"os"
	"path/filepath"
)

type DirFS string

var _ fs.FS = (*DirFS)(nil)

func (dir DirFS) Open(name string) (fs.File, error) {
	return os.Open(filepath.Join(string(dir), filepath.FromSlash(name)))
}

func (dir DirFS) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(filepath.Join(string(dir), filepath.FromSlash(name)))
}

func (dir DirFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(filepath.Join(string(dir), filepath.FromSlash(name)))
}

func (dir DirFS) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(string(dir), filepath.FromSlash(name)))
}

func (dir DirFS) Glob(pattern string) ([]string, error) {
	return filepath.Glob(filepath.Join(string(dir), pattern))
}
