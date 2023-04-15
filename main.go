package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Check if file path argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Please provide a path to a file containing the population of numbers")
		os.Exit(1)
	}

	// Open file and create scanner to read population
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	population := []float64{}

	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			os.Exit(1)
		}
		population = append(population, num)
	}

	// Calculate statistics
	average := calculateAverage(population)
	median := calculateMedian(population)
	variance := calculateVariance(population, average)
	stdDeviation := calculateStandardDeviation(variance)

	// Print results
	fmt.Println("Average:", round(average))
	fmt.Println("Median:", round(median))
	fmt.Println("Variance:", round(variance))
	fmt.Println("Standard Deviation:", round(stdDeviation))
}

func calculateAverage(population []float64) float64 {
	sum := 0.0
	for _, num := range population {
		sum += num
	}
	return sum / float64(len(population))
}

func calculateMedian(population []float64) float64 {
	sort.Float64s(population)
	length := len(population)

	if length%2 == 0 {
		return (population[length/2-1] + population[length/2]) / 2
	} else {
		return population[length/2]
	}
}

func calculateVariance(population []float64, mean float64) float64 {
	sum := 0.0
	for _, num := range population {
		sum += math.Pow(num-mean, 2)
	}
	return sum / float64(len(population))
}

func calculateStandardDeviation(variance float64) float64 {
	return math.Sqrt(variance)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
