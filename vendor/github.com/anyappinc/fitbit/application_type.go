package fitbit

// ApplicationType represents the type of registered application.
type ApplicationType int64

const (
	PersonalApplication ApplicationType = iota // PersonalApplication represents personal type application
	ClientApplication                          // ClientApplication represents client type application
	ServerApplication                          // ServerApplication represents server type application
)
