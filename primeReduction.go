package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

// NunoFromFile is used to read input data from input file (instead shell) and write results save it in output file
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

// Nuno is a classical approach to solve Kattis Challenge, reading input directly from os.Stdin
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

// isExitCode return true if number == 4
func isExitCode(number int64) bool {
	return number == 4
}

// primeReduction will print a string containg a prime number and a count with how many cycles was necessary to factor it.
// 1. If the value of the number is a prime number, print the result;
// 2. If the value of the number is not a prime number, start the factorization;
// 2.1. While the number is divisible by 2, divide it and add 2 to the sum variable for each successful division;
// 2.2. While the number is divisible by odd numbers, divide it and some divisor value to the sum variable for each successful division;
// 2.3. If the factorization result is a prime number, add it to the sum variable;
// 2.4. added 1 to the counter, substitute the value of the number received by the value of the variable sum and go back to step 1;
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

// factoringBy divides the number still it is divisible by factor and sums the factors used into variable 'sum'
func factoringBy(number *int64, sum *int64, factor int64) {
	for *number%factor == 0 {
		*number /= factor
		*sum += factor
	}
}

// factoringByOdds generate odd numbers (starting by 3) to call function factoringBy
func factoringByOdds(number *int64, sum *int64) {
	for i := int64(3); i*i <= *number; i += 2 {
		factoringBy(number, sum, i)
	}
}

// isPrimeNumber verify if a number is prime
func isPrimeNumber(number int64) bool {
	return big.NewInt(number).ProbablyPrime(0)
}
