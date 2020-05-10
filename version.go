package xjutils

import (
	"fmt"
)

// set on build time
var (
	GitCommit = "7b32838"
	BuildTime = "2020-05-10 12:29"
	GoVersion = "1.14"
	Version   = "0.0.1"
)

// PrintVersion Print out version information
func PrintVersion() {
	fmt.Println("Version  : ", Version)
	fmt.Println("GitCommit: ", GitCommit)
	fmt.Println("BuildTime: ", BuildTime)
	fmt.Println("GoVersion: ", GoVersion)
}
