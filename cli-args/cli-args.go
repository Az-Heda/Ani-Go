package cliargs

import (
	"os"
)

var allCliArgs map[string]interface{} = map[string]interface{}{
	"--help":    showHelp,
	"--init-db": initializeDB,
	"--import":  importFile,
}

func ParseArguments() {
	var count int = 0
	keys := make([]string, 0, len(allCliArgs))
	for k := range allCliArgs {
		keys = append(keys, k)
	}
	for _, arg := range os.Args[1:] {

		for k, v := range allCliArgs {
			if arg == k {
				v.(func(int, []string))(count, keys)
			}
		}
	}
}
