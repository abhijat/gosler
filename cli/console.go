package cli

import (
	"strings"
	"github.com/peterh/liner"
	"os"
	"log"
	"errors"
	"fmt"
)

var (
	commands        = [...]string{"set-prompt", "health-probe", "read-secret", "read-policy"}
	historyFilename = ".history"
)

func completionEngine(input string) []string {
	i := strings.ToLower(input)

	var candidates []string
	for _, command := range commands {
		if strings.HasPrefix(command, i) {
			candidates = append(candidates, command)
		}
	}

	return candidates
}

type console struct {
	*liner.State
	prompt string
}

func (c *console) Exit() {
	if file, err := os.Create(historyFilename); err != nil {
		log.Println("failed to write history to file:", err)
	} else {
		c.WriteHistory(file)
		file.Close()
	}

	os.Exit(0)
}

func (c *console) SwitchPrompt(s string) {
	if len(s) <= 11 {
		printResponse(nil, errors.New("set-prompt <prompt>"))
	} else {
		c.prompt = fmt.Sprintf("%s ", s[11:])
	}
}

func NewConsole(prompt string) *console {
	l := liner.NewLiner()

	l.SetTabCompletionStyle(liner.TabPrints)
	l.SetCompleter(completionEngine)

	loadHistory(l, historyFilename)

	return &console{l, prompt}
}

func loadHistory(liner *liner.State, filename string) {
	if file, err := os.Open(filename); err == nil {
		liner.ReadHistory(file)
		file.Close()
	}
}
