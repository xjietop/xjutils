package xjutils

import (
	"testing"

)

func Test_findByPk(t *testing.T) {
	GitCommit = "7b32839"
	BuildTime = "2020-05-10 12:27"
	GoVersion = "1.14"
	Version = "0.0.1"
	PrintVersion()
	t.Log(Version)
}
