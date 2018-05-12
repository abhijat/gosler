package cli

import (
	"github.com/fatih/color"
	"github.com/chzyer/readline"
	"encoding/json"
	"bytes"
)

func printOk(b bytes.Buffer) {
	color.New(color.FgGreen).Add(color.Bold).Print(string(b.Bytes()))
}

func printError(e error) {
	color.New(color.FgRed).Add(color.Bold).Println("error:", e)
}

func printResponse(b []byte, e error) {
	if e != nil {
		printError(e)
		return
	}

	var buf bytes.Buffer
	err := json.Indent(&buf, b, "", "  ")

	if err != nil {
		printError(err)
		return
	}

	printOk(buf)
}

func prompt(s string) string {
	return color.New(color.FgBlue).Add(color.Bold).Sprint(s)
}

func NewRL() (*readline.Instance, error) {
	return readline.NewEx(&readline.Config{
		Prompt:          prompt("[gosler] "),
		HistoryFile:     "./.history",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "good bye",
	})
}
