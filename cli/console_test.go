package cli

import (
	"testing"

	"bytes"
	"strings"
	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
	old := consoleOutput

	var b bytes.Buffer
	consoleOutput = &b

	f()

	consoleOutput = old
	return strings.TrimSpace(string(b.Bytes()))
}

func Test_completionEngine(t *testing.T) {
	candidates := completionEngine("r")
	assert.Equal(t, []string{"read-secret", "read-policy"}, candidates)
}

func Test_console_SwitchPrompt(t *testing.T) {
	c := NewConsole("")
	defer c.Close()

	output := captureOutput(func() {
		c.SwitchPrompt("set-prompt")
	})

	assert.Equal(t, "error: set-prompt <prompt>", output)

	c.SwitchPrompt("set-prompt [abcd]")
	assert.Equal(t, "[abcd] ", c.prompt)
}

func Test_loadHistory(t *testing.T) {
}
