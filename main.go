package main

import (
	"tbd/cmd"
)

func init() {
	cmd.SetUpFlags()
}

func main() {
	cmd.Execute()
}
