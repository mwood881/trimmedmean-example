# Trimmed Mean Example R vs Go

## Description
This example go code shows the trimmed mean in R and Go. Specifically, it uses the code from trimmean that I created as a Go package. Trimmed mean helps reduce outliers by reducing the percents on each end of the spectrum. 

## Features: 
- R script is included to show the R calculations of the mean function
- The Rscript helped determine the CSV files used in this example for Go so that the datsets were the same and the seeds were the same to identify if the results are comparable
- Go uses trimmean package
- Fair comparison by using integers.csv and floats.csv for both

## Setup:

(1) Installation


 ```bash
   git clone https://github.com/mwood881/trimmedmean-example.git
 ```


(2) for R: 
Run the script which should produce the integers and float csv files that are already documented. 
You can see that the R results are: 
INTEGERS Trimmed Mean in R: 488 
FLOATS Trimmed Mean in R: 498.7642

(3) For Go 
 ```bash
   go get github.com/mwood881/trimmean
 ```
install this package to make sure it is in the system


(4) ```bash
   go run main.go
 ```
You should get the following for Go:
Trimmed mean for integers: 488.00
Trimmed mean for floats: 498.76


## Results Comparison
Both R and Go resulted in 
Trimmed mean for integers: 488.00
Trimmed mean for floats: 498.76

Therefore, the results are consistent for both programming languages

Although R uses mean and trim arguments while Go uses the trimmean package, the results are the same.

## Resources used:
Github was used by asking how to make the csv files similar. It resulted in just doing this for both floats and integers floats <- runif(100, min = 1.0, max = 1000.0)  # Random floats between 1 and 1000
write.csv(floats, "floats.csv", row.names = FALSE). 




