# R program made the integer and float datasets

set.seed(123)  # Set seed for reproducibility
integers <- sample(1:1000, 100)
write.csv(integers, "integers.csv", row.names = FALSE)  # Save to file


set.seed(123)  # Set seed for reproducibility
floats <- runif(100, min = 1.0, max = 1000.0)  # Random floats between 1 and 1000
write.csv(floats, "floats.csv", row.names = FALSE)  # Save to file


# Compute the trimmed mean in R
integers_trimmed_mean_r <- mean(integers, trim = 0.05)
floats_trimmed_mean_r <- mean(floats, trim = 0.05)
cat("INTEGERS Trimmed Mean in R:", integers_trimmed_mean_r, "\n")
# INTEGERS Trimmed Mean in R: 488 
cat("FLOATS Trimmed Mean in R:", floats_trimmed_mean_r, "\n")
# FLOATS Trimmed Mean in R: 498.7642 
