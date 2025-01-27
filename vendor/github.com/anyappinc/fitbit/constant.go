package fitbit

const (
	apiBaseURL               = "https://api.fitbit.com"
	dateFormat               = "2006-01-02"                 // dateFormat is a format string to represent date
	CodeChallengeMethod      = "S256"                       // CodeChallengeMethod is the method used to hash the code challenge
	NumberLetters            = "0123456789"                 // NumberLetters is a set of characters represent numbers
	UppercaseAlphabetLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" // UppercaseAlphabetLetters is a set of upper case alphabetic characters
	LowercaseAlphabetLetters = "abcdefghijklmnopqrstuvwxyz" // LowercaseAlphabetLetters is a set of lower case alphabetic characters
)

var (
	apiEndpoints = map[string]string{
		"GetDailyActivitySummary": "/1/user/%s/activities/date/%s.json",
		"IntrospectToken":         "/1.1/oauth2/introspect",
		"RevokeToken":             "/oauth2/revoke",
		"GetWater":                "/1/user/%s/foods/log/water/date/%s.json",
		"GetProfile":              "/1/user/%s/profile.json",
	}
)
