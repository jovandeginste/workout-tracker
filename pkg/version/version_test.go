package version_test

import (
	"testing"

	"github.com/jovandeginste/workout-tracker/v2/pkg/version"
	"github.com/stretchr/testify/assert"
)

func TestVersion_PrettyVersion(t *testing.T) {
	v := version.Version{
		RefType: "branch",
		RefName: "master",
		Ref:     "master",
		Sha:     "abc123",
	}

	assert.Equal(t, v.PrettyVersion(), "branch master (abc123)")
}

func TestVersion_UserAgent(t *testing.T) {
	v := version.Version{
		RefType: "branch",
		RefName: "master",
		Ref:     "master",
		Sha:     "abc123",
	}

	assert.Equal(t, v.UserAgent(), "workout-tracker/master")
}
