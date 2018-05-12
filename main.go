package main

import (
	"github.com/chzyer/readline"
	"log"
	"os"
	"io"
	"strings"
	"strconv"
	"io/ioutil"
	"fmt"
)

var completer = readline.NewPrefixCompleter(
	readline.PcItem("set-prompt"),
	readline.PcItem("ls",
		readline.PcItem("files"),
		readline.PcItem("people"),
		readline.PcItem("images"),
		readline.PcItem("movies"),
		readline.PcItem("documents"),
	),
	readline.PcItem(
		"say",
		readline.PcItemDynamic(
			listFiles("./"),
			readline.PcItem("with",
				readline.PcItem("following"),
				readline.PcItem("items"),
			),
		),
	),
)

func listFiles(path string) func(string) []string {
	return func(line string) []string {
		fileNames := make([]string, 0)
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Println("unable to read files in path", path)
			return fileNames
		}

		for _, f := range files {
			fileNames = append(fileNames, f.Name())
		}

		return fileNames
	}
}

func makePrompt(s string) string {
	return fmt.Sprintf("\033[31m %s »\033[0m ", s)
}

func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          makePrompt("(gosler)"),
		HistoryFile:     "./.history",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "good bye",
	})

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

		case strings.HasPrefix(line, "set-prompt"):

			if len(line) <= 11 {
				log.Println("set-prompt <prompt>")
				break
			}

			l.SetPrompt(makePrompt(line[11:]))

		case line == "exit":
			os.Exit(0)

		default:
			log.Println("you said:", strconv.Quote(line))

		}
	}

}
