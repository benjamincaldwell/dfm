package main

import "github.com/benjamincaldwell/dfm/dfm"

var Version string
var BuildDate string

func main() {
	dfm.Version = Version
	dfm.BuildDate = BuildDate
	dfm.Execute()
}
