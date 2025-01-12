package caches

type task interface {
	GetId() string
	Run()
}
