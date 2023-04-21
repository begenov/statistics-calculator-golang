# statistics-calculator-golang

### Objectives

This project is a command-line tool that calculates various statistical measures for a population of numbers. The tool reads data from a file and calculates the following:

 -   Average
 -   Median
 -   Variance
 -   Standard Deviation


 ### Installation

 To use this tool, you need to have Go installed on your machine. Then, you can clone the repository and build the executable file by running the following commands in your terminal:

```
git clone https://github.com/begenov/statistics-calculator.git
cd statistics-calculator
go build
```

This will create an executable file called `statistics-calculator` in the current directory.

### Usage

To use the tool, you need to have a data file containing one number per line. The tool will calculate the statistics for the population of numbers in the file.

To run the tool, execute the following command:

```
./statistics-calculator path/to/data-file.txt
```

This will print the calculated statistics to the console. The output will be in the following format:

```
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```

### Contributing

If you find a bug or have a suggestion for a new feature, please open an issue in the repository. Pull requests are also welcome.
