package cli

import (
	"encoding/json"
	"bytes"
	"fmt"
)

func printOk(b bytes.Buffer) {
	fmt.Println(string(b.Bytes()))
}

func printError(e error) {
	fmt.Println("error:", e)
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
