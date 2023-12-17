package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input_file string

func main() {
	input_file = "./input.txt"
	file, err := os.ReadFile(input_file)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %v: %v\n", input_file, err))
	}

	input := string(file)
    // fmt.Println(input)
    final_value, err := get_calibration_value(strings.Split(input, "\n"))
    if err != nil {
        panic(err)
    }
    fmt.Printf("Final Total: %d\n", final_value)
}

func get_calibration_value(document []string) (int, error) {
	var total int

	for _, line := range document {
        if len(strings.TrimSpace(line)) == 0 {
            continue
        }
        var digits string
        var last byte
		for i := 0; i < len(line); i++ {
			if char := line[i]; char >= '0' && char <= '9' {
				if len(digits) == 0 {
                    digits += string(char)
				}
				last = char
			}
		}
        if len(digits) == 0 {
            return -1, fmt.Errorf("Invalid line, no digits found: %s\n", line)
        } else if last == 0 {
            digits += digits
        } else {
            digits += string(last)
        }
        // fmt.Printf("digits: %s\n", digits)
        number, err := strconv.Atoi(digits)
        if err != nil {
            return total, fmt.Errorf("Could not convert parsed digits to number: %s", digits)
        }
        total += number
        // fmt.Printf("Running Total: %d\n", total)
	}
	return total, nil
}
