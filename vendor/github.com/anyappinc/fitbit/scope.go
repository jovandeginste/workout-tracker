package fitbit

import (
	"strings"
)

// ScopeType represents the type of scope.
type ScopeType int64

const (
	// ScopeUnknown represents an unknown type of scope.
	ScopeUnknown ScopeType = iota
	// ScopeReadOnly means it is allowed to read data within the scope.
	ScopeReadOnly
	// ScopeReadWrite means it is allowed to read and write data within the scope.
	ScopeReadWrite
	// ScopeInvalid represents an invalid scope.
	ScopeInvalid
)

// MarshalJSON implements the json.Marshaler interface.
func (st ScopeType) MarshalJSON() ([]byte, error) {
	switch st {
	case ScopeReadOnly:
		return []byte(`"READ"`), nil
	case ScopeReadWrite:
		return []byte(`"READ_WRITE"`), nil
	case ScopeInvalid:
		return []byte(`"INVALID"`), nil
	}
	return []byte(`"UNKNOWN"`), nil
}

func convertToScopeType(s string) ScopeType {
	switch strings.ToLower(s) {
	case "read":
		return ScopeReadOnly
	case "read_write":
		return ScopeReadWrite
	}
	return ScopeUnknown
}

// Scope represents the scope of permission.
type Scope struct {
	Activity  bool
	Heartrate bool
	Location  bool
	Nutrition bool
	Profile   bool
	Settings  bool
	Sleep     bool
	Social    bool
	Weight    bool
}

func newScope(raw []string) *Scope {
	scope := &Scope{}
	for _, s := range raw {
		switch strings.ToLower(s) {
		case "activity":
			scope.Activity = true
		case "heartrate":
			scope.Heartrate = true
		case "location":
			scope.Location = true
		case "nutrition":
			scope.Nutrition = true
		case "profile":
			scope.Profile = true
		case "settings":
			scope.Settings = true
		case "sleep":
			scope.Sleep = true
		case "social":
			scope.Social = true
		case "weight":
			scope.Weight = true
		}
	}
	return scope
}

func parseScopeFromTokenState(scopeString string) (*Scope, ScopeType) {
	scopeString = strings.TrimPrefix(scopeString, "{")
	scopeString = strings.TrimSuffix(scopeString, "}")
	scopeParts := strings.Split(scopeString, ", ")
	if scopeParts[0] == "" {
		return &Scope{}, ScopeInvalid
	}
	rawScope := make([]string, len(scopeParts))
	for i, part := range scopeParts {
		rawScope[i] = strings.SplitN(part, "=", 2)[0]
	}
	return newScope(rawScope), convertToScopeType(strings.SplitN(scopeParts[0], "=", 2)[1])
}

func (s *Scope) convert() []string {
	scopes := make([]string, 0, 9)
	if s.Activity {
		scopes = append(scopes, "activity")
	}
	if s.Heartrate {
		scopes = append(scopes, "heartrate")
	}
	if s.Location {
		scopes = append(scopes, "location")
	}
	if s.Nutrition {
		scopes = append(scopes, "nutrition")
	}
	if s.Profile {
		scopes = append(scopes, "profile")
	}
	if s.Settings {
		scopes = append(scopes, "settings")
	}
	if s.Sleep {
		scopes = append(scopes, "sleep")
	}
	if s.Social {
		scopes = append(scopes, "social")
	}
	if s.Weight {
		scopes = append(scopes, "weight")
	}
	return scopes
}

// Missing returns a list of missing scope as a string slice.
func (s *Scope) Missing(expected *Scope) []string {
	missingScopes := make([]string, 0, 9)
	if expected.Activity && !s.Activity {
		missingScopes = append(missingScopes, "activity")
	}
	if expected.Heartrate && !s.Heartrate {
		missingScopes = append(missingScopes, "heartrate")
	}
	if expected.Location && !s.Location {
		missingScopes = append(missingScopes, "location")
	}
	if expected.Nutrition && !s.Nutrition {
		missingScopes = append(missingScopes, "nutrition")
	}
	if expected.Profile && !s.Profile {
		missingScopes = append(missingScopes, "profile")
	}
	if expected.Settings && !s.Settings {
		missingScopes = append(missingScopes, "settings")
	}
	if expected.Sleep && !s.Sleep {
		missingScopes = append(missingScopes, "sleep")
	}
	if expected.Social && !s.Social {
		missingScopes = append(missingScopes, "social")
	}
	if expected.Weight && !s.Weight {
		missingScopes = append(missingScopes, "weight")
	}
	return missingScopes
}
