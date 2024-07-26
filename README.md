# Go Calculator

A simple command-line calculator that supports both Arabic and Roman numerals for basic arithmetic operations (+, -, *, /).

## Features

- Supports addition, subtraction, multiplication, and division
- Handles both Arabic (1, 2, 3, ...) and Roman (I, II, III, ...) numerals
- Input numbers from 1 to 10 inclusively
- Outputs results in the same numeral system as the input
- Performs integer division (remainder is discarded)

## Usage

### Prerequisites

- [Go](https://golang.org/dl/) installed

### Running the Calculator

1. Clone the repository:

    ```sh
    git clone git@github.com:lelonov23/KataCalculator.git
    cd KataCalculator
    ```

2. Run the application:

    ```sh
    go run main.go
    ```

3. Follow the prompts to enter your arithmetic expression. The expression must be in the format:
    
    ```
    <number1> <operator> <number2>
    ```

    Example inputs:
    
    - `1 + 2`
    - `VI / III`
    - `X - II`
    - `4 * 3`

### Example

```sh
$ go run main.go
Enter expression:
3 + 4
Result: 7

$ go run main.go
Enter expression:
V * IV
Result: XX
