package cli

import (
	"github.com/chzyer/readline"
	"io/ioutil"
	"log"
)

var completer *readline.PrefixCompleter

func init() {
	completer = readline.NewPrefixCompleter(
		readline.PcItem("set-prompt"),
		readline.PcItem("health-probe"),
		readline.PcItem("read-secret"),
	)
}

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

