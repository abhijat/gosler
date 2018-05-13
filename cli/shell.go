package cli

import (
	"io"
	"strings"
	"github.com/abhijat/gosler/vault"
	"errors"
	"fmt"
)

func Shell(client *vault.Client) {

	console := NewConsole("[gosler] ")
	defer console.Close()

	for {
		line, err := console.Prompt(console.prompt)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("failed to read input:", err)
			break
		}

		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		console.AppendHistory(line)

		if strings.HasPrefix(line, "set-prompt") {
			console.SwitchPrompt(line)
		}

		if line == "exit" {
			console.Exit()
		}

		if line == "health-probe" {
			printResponse(client.HealthProbe())
		}

		if strings.HasPrefix(line, "read-secret") {
			fields := strings.Fields(line)
			if len(fields) != 2 {
				printResponse(nil, errors.New("need a secret path with read-secret"))
				continue
			}
			printResponse(client.ReadSecret(fields[1]))
		}
	}
}
