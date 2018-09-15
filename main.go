package main

import (
	"log"

	"github.com/jkamenik/log-pruner/cmd"
)

func main() {
	cmd.Execute()
}

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
}
