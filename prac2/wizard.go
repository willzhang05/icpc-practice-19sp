package main

import (
    "os"
    "bufio"
    "fmt"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        n, _ := strconv.Atoi(line)
        for i:=1; i <= n; i++ {
            fmt.Print(i)
            fmt.Println(" Abracadabra")
        }
    }
}
