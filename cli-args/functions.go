package cliargs

import (
	dbintegration "AniGo/db-integration"
	"fmt"
	"os"
)

func showHelp(ec int, keys []string) {
	fmt.Println("CLI Args list:")
	for _, key := range keys {
		fmt.Printf("%s\n", key)
	}
	os.Exit(ec)
}

func initializeDB(ec int) {
	dbintegration.Init()
	os.Exit(ec)
}
