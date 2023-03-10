package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	createFile("me.txt", "hello from golang dev")
	var res []string
	readFile("mm.txt", &res)
	fmt.Println(res)
}

func createFile(filename string, content string) {
	_, err := os.Stat(filename)
	if err == nil {
		fmt.Println("This filename exist")
	}
	fileCreated, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileCreated.WriteString(content)
	defer fileCreated.Close()
}

func readFile(filename string, content *[]string) {
	openedFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer openedFile.Close()
	scanner := bufio.NewScanner(openedFile)
	for scanner.Scan() {
		*content = append(*content, scanner.Text())
	}
}
