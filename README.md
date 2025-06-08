# Largest Number from Concatenation

This is a Go implementation of a coding challenge that takes a list of non-negative integers and rearranges them to form the **largest possible number** when concatenated.

## Problem Statement

Given a slice of non-negative integers, rearrange them such that they form the largest number.

### Examples

#### Example 1
```
Input: []int{10, 2}
Output: "210"
```

#### Example 2
```
Input: []int{3, 80, 5, 9}
Output: "98053"
```

## How to Run

1. Make sure you have Go 1.18+ installed.
2. Clone this repository:
    ```
    git clone git@github.com:mghifariyusuf/largest-number-go.git
    cd largest-number-go
    ```
3. Run the program:
    ```
    go run main.go
    ```

## Running Tests
This project includes unit tests.

To run tests:
```
go test
```


## File Structure

- main.go – Contains the main function and the LargestNumber logic.
- main_test.go – Unit tests for the logic.

## Algorithm Overview

The main idea is to sort the numbers as strings, using a custom comparison rule:

- For two strings a and b, compare a + b and b + a.
- If a + b > b + a, then a should come before b.

This ensures the final concatenated result is the largest possible.

Edge case handled:
- If all values are 0 (e.g., [0, 0]), the output will be "0".

## License

This project is licensed under the MIT License.