package cli

import (
	"encoding/json"
	"bytes"
	"fmt"
	"os"
	"io"
	"strings"
)

var consoleOutput io.Writer = os.Stdout

func formatSuccess(b bytes.Buffer) string {
	return string(b.Bytes())
}

func formatError(e error) string {
	return fmt.Sprint("error: ", e)
}

func formatResponse(b []byte, e error) string {
	if e != nil {
		return formatError(e)
	}

	var buf bytes.Buffer
	err := json.Indent(&buf, b, "", "  ")

	if err != nil {
		return formatError(err)
	}

	return formatSuccess(buf)
}

func printResponse(b []byte, e error) {
	fmt.Fprintln(consoleOutput, formatResponse(b, e))
}

func captureConsoleOutput(f func()) string {
	old := consoleOutput

	var b bytes.Buffer
	consoleOutput = &b

	f()

	consoleOutput = old
	return strings.TrimSpace(string(b.Bytes()))
}
