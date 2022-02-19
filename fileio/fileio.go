package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	testfilename = "testfile.txt"
)

func main() {
	// get current working directory
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.Open(path + "/" + testfilename)
	check(err)

        // closure of file
        defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for _, each_ln := range text {
		fmt.Println(each_ln)
	}
}
