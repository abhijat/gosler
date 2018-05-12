package cli

import (
	"log"
	"os"
	"io"
	"strings"
	"github.com/abhijat/gosler/vault"
	"errors"
	"github.com/chzyer/readline"
)

func Shell(client *vault.Client) {
	l, err := NewRL()
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	log.SetOutput(os.Stderr)

	for {
		line, err := l.Readline()
		if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)

		switch {

		case len(line) == 0:
			continue

		case strings.HasPrefix(line, "set-prompt"):
			switchPrompt(line, l)

		case line == "exit":
			os.Exit(0)

		case line == "health-probe":
			printResponse(client.HealthProbe())

		case strings.HasPrefix(line, "read-secret "):
			fields := strings.Fields(line)
			if len(fields) != 2 {
				printResponse(nil, errors.New("need a secret path with read-secret"))
				continue
			}

			printResponse(client.ReadSecret(fields[1]))
		}
	}
}

func switchPrompt(s string, l *readline.Instance) {
	if len(s) <= 11 {
		printResponse(nil, errors.New("set-prompt <prompt>"))
	} else {
		l.SetPrompt(prompt(s[11:]))
	}
}
