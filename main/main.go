// https://www.hackerrank.com/challenges/security-function-ii/problem
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Printf("Failed to parse the input. %v.", err.Error())
            return
        }

        fmt.Println(num * num)
    }
}
