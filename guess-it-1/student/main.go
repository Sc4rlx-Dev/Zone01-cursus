package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var previousnumbers []float64
	for scanner.Scan() {
		numba9istring:=scanner.Text()
		number, err := strconv.ParseFloat(numba9istring,64)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		previousnumbers = append(previousnumbers, number)
		if len(previousnumbers) > 1 {
			lower, upper := guessit(previousnumbers)
			fmt.Printf("%d %d\n", lower, upper)
		} else {
			lowerf,up := (number-10),(number+10)
			fmt.Printf("%d %d\n", int(math.Round(lowerf)), int(math.Round(up)))	
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}
func guessit(numbers []float64) (int, int) {
	var (
		mean     float64
		sd       float64
		sum      float64
		variance float64
	)

	n := float64(len(numbers))
	for _, v := range numbers {
		sum += v
	}
	mean = sum / n 
	for _, v := range numbers {
		variance += ((v - mean) * (v - mean))
	}
	sd = math.Sqrt(variance / n) 

	lowerLimit := int(math.Round(mean - (2*sd)))
	upperLimit := int(math.Round(mean + (2*sd)))
	return lowerLimit, upperLimit
}

// package main

// import (
//     "bufio"
//     "fmt"
//     "os"
//     "strconv"
// )

// func main() {
//     scanner := bufio.NewScanner(os.Stdin)
//     var previous int
//     var isFirstInput = true

//     for scanner.Scan() {
//         // Read the current input
//         currentInput := scanner.Text()
//         current, err := strconv.Atoi(currentInput)
//         if err != nil {
//             fmt.Println("Invalid input. Please enter a valid number.")
//             continue
//         }

//         // Print the current input
//         fmt.Println(current)

//         // If it's not the first input, calculate the range
//         if !isFirstInput {
//             // Calculate the range
//             lowerLimit := previous - 10
//             upperLimit := previous + 10

//             // Print the range
//             fmt.Printf("%d %d\n", lowerLimit, upperLimit)
//         }

//         // Update the previous number
//         previous = current
//         isFirstInput = false
//     }

//     if err := scanner.Err(); err != nil {
//         fmt.Fprintln(os.Stderr, "Error reading input:", err)
//     }
// }