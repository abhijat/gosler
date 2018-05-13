package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_completionEngine(t *testing.T) {

	testCases := []struct {
		input      string
		candidates []string
	}{
		{input: "r", candidates: []string{"read-secret", "read-policy"}},
		{input: "read-x", candidates: nil},
		{input: "set", candidates: []string{"set-prompt"}},
	}

	for _, testCase := range testCases {
		candidates := completionEngine(testCase.input)
		assert.Equal(t, testCase.candidates, candidates, fmt.Sprintf("result does not match for input: %s", testCase.input))
	}
}

func Test_console_SwitchPrompt(t *testing.T) {
	c := NewConsole("")
	defer c.Close()

	output := captureConsoleOutput(func() {
		c.SwitchPrompt("set-prompt")
	})

	assert.Equal(t, "error: set-prompt <prompt>", output)

	c.SwitchPrompt("set-prompt [abcd]")
	assert.Equal(t, "[abcd] ", c.prompt)
}
