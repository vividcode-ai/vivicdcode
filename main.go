package main

import (
	"github.com/vividcode-ai/vividcode/cmd"
	"github.com/vividcode-ai/vividcode/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
