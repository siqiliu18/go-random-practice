# go-random-practice

## Note: 2023-11-10
1. The diff between new and make in Go: 
    - new just create a pointer for the data structure
    - make create a pointer for the data structure AND also init it by providing default values for it
2. Find fibonacci and prime values in Go
    - we can use recurrsion to find fibonacci but duplicated values will be calculated. If more memory usage is acceptable, we bettwe use bottom up dynamic programming for it.
    - to find prime, Go has a built-in function already https://yourbasic.org/golang/check-prime/ and it says "This primality test is 100% accurate for inputs less than 264."