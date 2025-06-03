package main

import (
	"log"
	"os"

	"slackstatus/cmd"
)

func main() {
	if err := cmd.Run(os.Args); err != nil {
		log.Fatalf("❌ Error: %v\n", err)
	}
}
