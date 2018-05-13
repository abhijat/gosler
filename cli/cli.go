package cli

import (
	"encoding/json"
	"bytes"
	"fmt"
)

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
	fmt.Println(formatResponse(b, e))
}
