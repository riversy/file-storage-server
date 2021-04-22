package main

import (
	"github.com/riversy/file-storage-server/cmd/server"
	"github.com/teris-io/cli"
	"os"
)

func DefaultAction(args []string, options map[string]string) int {
	err := server.RunGrpcServer()
	if err != nil {
		return 1
	}
	return 0
}

func main() {
	app := cli.New("gRPC Files Service").
		WithAction(DefaultAction)

	os.Exit(app.Run(os.Args, os.Stdout))
}
