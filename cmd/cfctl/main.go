/*
Copyright © 2022 jaronnie jaron@jaronnie.com

*/

package main

import (
	"github.com/jaronnie/cfc/cmd/cfctl/internal/cmd"
)

var version string

func main() {
	cmd.Version = version
	cmd.Execute()
}
