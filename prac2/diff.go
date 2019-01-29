package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        nums := strings.Split(line, " ")
        a, _ := strconv.Atoi(nums[0])
        b, _ := strconv.Atoi(nums[1])
        if (a > b) {
            fmt.Println(a-b)
        } else {
            fmt.Println(b-a)
        }
    }

}
