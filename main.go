package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	prefixToIgnore := "/some/long/path"
	currentDir := os.Args[1]
	split := strings.Split(currentDir, "/")
	finalDir := split[len(split)-1]
	if strings.HasPrefix(currentDir, prefixToIgnore) {
		fmt.Printf("\012\\[\\e[0;35m\\]\\t \\[\\e[0;32m\\]\\u@\\h \\[\\e[0;33m\\]%s\\[\\e[m\\]\\n$", finalDir)
	} else {
		fmt.Print("kek")
	}
}
