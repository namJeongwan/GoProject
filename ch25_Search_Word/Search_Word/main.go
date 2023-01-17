package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type stringArray []string

func (strArr *stringArray) String() string {
	return "target Files.."
}

func (strArr *stringArray) Set(value string) error {
	*strArr = append(*strArr, value)
	return nil
}

const DEFAULT_VALUE = ""

func main() {
	wordArg := flag.String("word", DEFAULT_VALUE, "target word")
	flag.Parse()

	fileArg := flag.Args()

	if *wordArg == DEFAULT_VALUE {
		panic("word args is required..")
	}

	fmt.Printf("Target word => %v\n", wordArg)

	targetFiles := PrintAndGetAllFiles(fileArg)

	fmt.Printf("### [Print Target Lines..] ###\n\n")
	for _, path := range targetFiles {
		FindAndPrintTargetFile(path, *wordArg)
	}

}

func FindAndPrintTargetFile(filename, word string) {
	fileObject, err := os.Open(filename) // 파일 읽기
	if err != nil {
		fmt.Printf("[%v] Can not open file, err: %v\n", filename, err)
		return
	}
	defer fileObject.Close()

	scanner := bufio.NewScanner(fileObject)
	n := 0
	for scanner.Scan() {
		n++
		line := scanner.Text()
		if strings.Contains(line, word) {
			fmt.Printf("[%v - %d] %v\n", filename, n, line)
		}
	}
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAndGetAllFiles(files []string) []string {
	var result []string
	for _, path := range files {
		fileList, err := GetFileList(path)
		if err != nil {
			fmt.Println("file is not exists.., err: ", err)
			return nil
		}
		fmt.Printf("### [File List => %v] ###\n", path)
		for _, name := range fileList {
			fmt.Println(name)
			result = append(result, name)
		}
		fmt.Println("")
	}
	return result
}
