package main

import "github.com/ttto/mergectl/cmd"

var version string

func main() {
	cmd.Version = version
	cmd.Execute()
}
