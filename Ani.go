package main

import (
	"os"

	cliargs "AniGo/cli-args"
	router "AniGo/router"
)

func main() {
	if len(os.Args) > 1 {
		cliargs.ParseArguments()
	}
	router.Init()
}
