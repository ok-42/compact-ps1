package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	prefixToIgnore := "/some/long/path"
	lastIndex := strings.LastIndex(prefixToIgnore, "/")
	prefixToIgnoreD := prefixToIgnore[:lastIndex+1]
	currentDir := os.Args[1]

	var result string = "\012"

	// Add venv to PS1
	venvName := os.Getenv("VIRTUAL_ENV")
	var projectName string
	if venvName != "" {
		venvNameSplit := strings.Split(venvName, "\\")
		projectName = venvNameSplit[(len(venvNameSplit) - 2)]
		result += fmt.Sprintf("\\[\\e[1;34m\\](%s) ", projectName)
	}

	// Add current dir to PS1
	var pathPlaceholder string
	if strings.HasPrefix(currentDir, prefixToIgnore) {
		if currentDir == prefixToIgnore {
			pathPlaceholder = prefixToIgnore[lastIndex+1:]
		} else {
			pathPlaceholder, _ = strings.CutPrefix(currentDir, prefixToIgnoreD)
		}
	} else {
		pathPlaceholder = "\\w"
	}

	// Print PS1           purple       time  green        user@host orange        pwd    erase colour
	result += fmt.Sprintf("\\[\\e[0;35m\\]\\t \\[\\e[0;32m\\]\\u@\\h \\[\\e[0;33m\\]%s\\[\\e[m\\]\\n$", pathPlaceholder)
	fmt.Println(result)
}
