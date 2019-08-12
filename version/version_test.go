package version

import (
	"fmt"
	"runtime"
	"testing"
)

func TestPrint(t *testing.T) {
	Version = "1.0.0"
	Revision = "4819e09c01edbb4d8bf019765b4d65cd254f34a5"
	Branch = "master"
	BuildUser = "ricoberger"
	BuildDate = "2019-01-01@12:00:00"

	_, err := Print("myapp")
	if err != nil {
		t.Errorf("Could not print version information: %s", err.Error())
	}
}

func TestInfo(t *testing.T) {
	Version = "1.0.0"
	Revision = "4819e09c01edbb4d8bf019765b4d65cd254f34a5"
	Branch = "master"
	BuildUser = "ricoberger"
	BuildDate = "2019-01-01@12:00:00"

	if Info() != "(version=1.0.0, branch=master, revision=4819e09c01edbb4d8bf019765b4d65cd254f34a5)" {
		t.Errorf("Bad formation for Info()")
	}
}

func TestBuildContext(t *testing.T) {
	Version = "1.0.0"
	Revision = "4819e09c01edbb4d8bf019765b4d65cd254f34a5"
	Branch = "master"
	BuildUser = "ricoberger"
	BuildDate = "2019-01-01@12:00:00"

	if BuildContext() != fmt.Sprintf("(go=%s, user=ricoberger, date=2019-01-01@12:00:00)", runtime.Version()) {
		t.Errorf("Bad formation for BuildContext()")
	}
}
