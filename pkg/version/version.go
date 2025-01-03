package version

import "fmt"

type Version struct {
	BuildTime string
	Ref       string
	RefName   string
	RefType   string
	Sha       string
}

func (v Version) PrettyVersion() string {
	rn := v.RefName

	if v.RefType == "branch" {
		rn = "branch " + rn
	}

	return fmt.Sprintf("%s (%.8s)", rn, v.Sha)
}

func (v Version) UserAgent() string {
	return "workout-tracker/" + v.Ref
}
