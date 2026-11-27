package partials

// RadioPillOption is one option in a RadioPillGroup.
type RadioPillOption struct {
	Value    string // The value submitted with the form
	LabelKey string // The i18n key for the visible label
	Checked  bool   // Whether this option is selected
}
