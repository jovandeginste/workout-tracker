package fit

// Logger mimics a subset of golang's standard Logger as an interface.
type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}
