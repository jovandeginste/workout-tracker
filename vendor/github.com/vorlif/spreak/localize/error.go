package localize

// Error is returned when a Localizer or Locale translates an error into the target language.
type Error struct {
	Translation string
	Wrapped     error
}

func (e *Error) Error() string {
	return e.Translation
}

func (e *Error) Unwrap() error {
	return e.Wrapped
}

func (e *Error) String() string {
	return e.Translation
}
