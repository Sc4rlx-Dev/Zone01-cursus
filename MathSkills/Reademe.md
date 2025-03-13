# Statistical Calculator

This program reads a list of integers from a file, computes statistical values, and displays the results.

## Features
- Computes **mean (average)**
- Computes **median**
- Computes **variance**
- Computes **standard deviation**

## Usage
Compile the program using GCC:
```sh

```

Run the program:
```sh
./stats_calculator <file_path>
```

## Input File Format
- The input file should contain **one integer per line**.
- Example:
  ```
  10
  20
  30
  40
  ```

## Example Output
```sh
Average: 25
Median: 25
Variance: 125
Standard Deviation: 11
```

## Error Handling
- Displays an error if the file is missing or unreadable.
- Displays an error if the file contains no valid numbers.

## License
This project is licensed under the MIT License.