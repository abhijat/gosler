package main

import (
	"github.com/abhijat/gosler/cli"
	"github.com/abhijat/gosler/vault"
	"os"
	"fmt"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <token>\n", os.Args[0])
		os.Exit(1)
	}
}

func main() {
	client := vault.NewClient(
		"http://localhost:8200/v1",
		os.Args[1],
	)

	cli.Shell(client)
}
