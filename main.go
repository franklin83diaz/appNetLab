package main

import (
	"appnetlab/cmd"
	"appnetlab/internal"
)

func main() {
	internal.CheckDependencies()
	cmd.Execute()
}
