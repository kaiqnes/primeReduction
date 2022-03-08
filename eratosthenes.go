package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Eratosthenes() {
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

		num := int(number)
		reduction := eratosthenesPrimeReduction(num, 1, classifyNumbers(num))
		fmt.Print(reduction)
	}
}

func EratosthenesFromFile(input, output string) {
	inputFile := getFile(input)
	outputFile := getFile(output)

	defer inputFile.Close()
	defer outputFile.Close()

	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	inputData := readFromFileToIntSlice(reader)

	for _, number := range inputData {
		if isExitCode(int64(number)) {
			break
		}
		result := eratosthenesPrimeReduction(number, 1, classifyNumbers(number))
		writeInResponseFile(result, writer)
	}
}

func eratosthenesPrimeReduction(currentNumber, count int, numbersClassification []bool) string {
	if isPrime(currentNumber, numbersClassification) {
		return fmt.Sprintf("%d %d\n", currentNumber, count)
	}

	factorSum := 0
	currentPrime := 2

	for {
		if currentNumber%currentPrime == 0 {
			currentNumber /= currentPrime
			factorSum += currentPrime

			if isPrime(currentNumber, numbersClassification) {
				factorSum += currentNumber
				break
			}
		} else {
			currentPrime = getNextPrime(currentPrime, numbersClassification)
		}
	}

	count++
	return eratosthenesPrimeReduction(factorSum, count, numbersClassification)
}

func classifyNumbers(limit int) []bool {
	var primeClassification []bool

	primeClassification = eratosthenesSieve(limit)

	return primeClassification
}

func segmentedEratosthenesSieve(limit int) []bool {
	// TODO: implement a Segmented Eratosthenes Sieve
	return []bool{}
}

func incrementalEratosthenesSieve(limit int) []bool {
	// TODO: implement an Incremental Eratosthenes Sieve
	return []bool{}
}

func eratosthenesSieve(limit int) []bool {
	var (
		integerSqrtLimit    = int(math.Sqrt(float64(limit)))
		primeClassification = make([]bool, limit+1)
	)

	// Fill even numbers
	for j := 4; j <= limit; j += 2 {
		primeClassification[j] = true
	}

	// Fill odd not primes
	for i := 3; i <= integerSqrtLimit; i += 2 {
		if !primeClassification[i] {
			for j := i * i; j <= limit; j += i {
				primeClassification[j] = true
			}
		}
	}

	return primeClassification
}

func isPrime(number int, classification []bool) bool {
	return !classification[number]
}

func getNextPrime(prime int, numbersClassification []bool) int {
	if prime%2 == 0 {
		return 3
	}
	for i := prime + 2; i < len(numbersClassification); i += 2 {
		if !numbersClassification[i] {
			return i
		}
	}
	return 0
}
