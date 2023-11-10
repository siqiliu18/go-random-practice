# go-random-practice

## Note: 2023-11-10
1. The diff between new and make in Go: 
    - *new* just create a pointer for the data structure
    - *make* create a pointer for the data structure AND also init it by providing default values for it
      * for creating a matrix or table, if we want all boxes within the matrix or table to have default values (0 for [][]int64 as an example), we need to use make for initiatizing the matrix or table first, then for the slice in each row, we also need to use make to init it as well. **(use make twice, one for the whole matrix or table, then for each row)**
      * **if we need to append values to each row, then don't need to use make to init each row, just append directly.** For example, rotate right for 90 degrees case. And for storing all nodes from each level of a tree. 
2. Find fibonacci and prime values in Go
    - we can use recurrsion to find fibonacci but duplicated values will be calculated. If more memory usage is acceptable, we bettwe use bottom up dynamic programming for it.
    - to find prime, Go has a built-in function already https://yourbasic.org/golang/check-prime/ and it says "This primality test is 100% accurate for inputs less than 264."
