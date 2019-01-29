package main

import (
    "os"
    "bufio"
    "fmt"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var s []rune= make([]rune, 0)
    for scanner.Scan() {
        line := scanner.Text()
        for _, char := range line {
            if (char == '(') {
                s = append(s, char)
            }
            if (char == ')') {
                if (len(s) != 0) { 
                    last := s[len(s)-1]
                    if (last == ')') {
                        fmt.Println("impossible")
                        os.Exit(0)
                    }
                    s = s[:len(s)-1]
                }
            }
        }
    }
}
