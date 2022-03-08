package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func NunoFromFile(input, output string) {
	inputFile := getFile(input)
	outputFile := getFile(output)

	defer inputFile.Close()
	defer outputFile.Close()

	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	inputData := readFromFileToInt64Slice(reader)

	for _, data := range inputData {
		if isExitCode(data) {
			break
		}
		result := primeReduction(data, 1)
		writeInResponseFile(result, writer)
	}
}

func Nuno() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		number, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if isExitCode(number) {
			break
		}

		reduction := primeReduction(number, 1)
		fmt.Println(reduction)
	}
}

func isExitCode(number int64) bool {
	return number == 4
}

func primeReduction(number int64, count int) string {
	if !isPrimeNumber(number) {
		var sum int64

		factoringBy(&number, &sum, 2)
		factoringByOdds(&number, &sum)

		if number > 2 {
			sum += number
		}

		return primeReduction(sum, count+1)
	}

	return fmt.Sprintf("%d %d\n", number, count)
}

func factoringBy(number *int64, sum *int64, factor int64) {
	for *number%factor == 0 {
		*number /= factor
		*sum += factor
	}
}

func factoringByOdds(number *int64, sum *int64) {
	for i := int64(3); i*i <= *number; i += 2 {
		factoringBy(number, sum, i)
	}
}

func isPrimeNumber(number int64) bool {
	return big.NewInt(number).ProbablyPrime(0)
}
