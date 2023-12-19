package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input_file string

func main() {
	// input_file = "./example_input_p2.txt"
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
	valid_string_digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	total := 0

	for _, line := range document {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		line = strings.ToLower(line)
		var digits string

		first_match_index := -1
		last_match_index := -1

		var first_digit_index, last_digit_index int

		for digit_index, digit := range valid_string_digits {
			// Check first occurrence
			match_index := strings.Index(line, digit)
			if match_index == -1 {
				continue
			}

			// fmt.Println("Found a match for", digit, "at index", match_index)

			if first_match_index == -1 {
				first_match_index = match_index
				first_digit_index = digit_index
			} else if match_index < first_match_index {
				first_match_index = match_index
				first_digit_index = digit_index
			}

			// Check last occurrence
			match_index = strings.LastIndex(line, digit)
			// fmt.Println("Found a LastIndex match for", digit, "at index", match_index)

			if match_index > last_match_index {
				last_match_index = match_index
				last_digit_index = digit_index
			}
		}

		if first_digit_index == -1 {
			return -1, fmt.Errorf("Invalid line, no digits found: %s\n", line)
		}

		if last_digit_index == -1 {
			return -1, fmt.Errorf("Invalid line, second digit not identified. (should duplicate the first digit: %s\n", line)
		}

		// fmt.Printf("line: %s\tfirst: %s\tlast: %s\n", line, valid_string_digits[first_digit_index], valid_string_digits[last_digit_index])
		// fmt.Println(first_digit_index, last_digit_index)

		if first_digit_index <= 9 {
			digits = fmt.Sprint(first_digit_index)
		} else {
			digits = valid_string_digits[first_digit_index]
		}

		if last_digit_index <= 9 {
			digits += fmt.Sprint(last_digit_index)
		} else {
			digits += valid_string_digits[last_digit_index]
		}

		// fmt.Printf("digits: %s\n", digits)

		number, err := strconv.Atoi(digits)
		if err != nil {
			return total, fmt.Errorf("Could not convert parsed digits to number: %s", digits)
		}

		// fmt.Printf("converted digits: %d\n", number)

		if digits != fmt.Sprint(number) {
			return total, fmt.Errorf("Error in ATOI converstion, number is not the same")
		}
		total += number
		// fmt.Printf("Running Total: %d\n", total)
	}
	return total, nil
}
