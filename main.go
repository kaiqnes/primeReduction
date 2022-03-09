package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	var1 = "sample.in"
)

func main() {
	executeFuncWithTimeTrack("Nuno", NunoFromFile)
	executeFuncWithTimeTrack("Eratosthenes", EratosthenesFromFile)
}

func executeFuncWithTimeTrack(name string, function func(string, string)) {
	init := time.Now()
	function(var1, fmt.Sprintf("%s.out", name))
	fmt.Printf("Func %s took %s\n", name, time.Since(init).String())
}

func writeInResponseFile(result string, output *bufio.Writer) {
	_, err := output.WriteString(result)
	if err != nil {
		panic(err)
	}
	output.Flush()
}

func readFromFileToIntSlice(input *bufio.Reader) []int {
	result := []int{}
	for {
		line, _, err := input.ReadLine()
		if line == nil {
			break
		}
		if err != nil {
			panic(err)
		}
		intResult, err := strconv.ParseInt(string(line), 10, 64)
		if err != nil {
			panic(err)
		}
		result = append(result, int(intResult))
	}

	return result
}

func readFromFileToInt64Slice(input *bufio.Reader) []int64 {
	result := []int64{}
	for {
		line, _, err := input.ReadLine()
		if line == nil {
			break
		}
		if err != nil {
			panic(err)
		}
		intResult, err := strconv.ParseInt(string(line), 10, 64)
		if err != nil {
			panic(err)
		}
		result = append(result, intResult)
	}

	return result
}

func getFile(fileName string) *os.File {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	return file
}
