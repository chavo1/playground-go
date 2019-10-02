Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression. 
Compare the performance of the two versions. (Section 11.4 shows how to compare the performance of different implementations systematically.)

Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument through 64 bit positions, testing the right most bit each time. Compare its performance to the table lookup version.

Exercise 2.5: The expression x&(x-1) clears the right most non-zero bit of x. Write a version of PopCount that counts bits by using this fact, and assess its performance.

