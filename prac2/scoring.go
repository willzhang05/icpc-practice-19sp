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
    var problems map[string]bool = make(map[string]bool)
    var attempts map[string]int = make(map[string]int)
    var solved []string = make([]string, 0)
    score := 0
    for scanner.Scan() {
        line := scanner.Text()
        if (line == "-1") {
            break
        }
        entry := strings.Split(line, " ")
        m, _ := strconv.Atoi(entry[0])
        prob := strings.TrimSpace(entry[1])
        status := entry[2]
        if (status == "right") {
            score += m
            solved = append(solved, prob)
            _, ok := problems[prob]
            if (!ok) {
                problems[prob] = true
            }
        } else {
            _, ok := attempts[prob]
            if (ok) {
                attempts[prob] += 1
            } else {
                attempts[prob] = 1
            }

            _, ok = problems[prob]
            if (!ok) {
                problems[prob] = false
            }
        }
    }
    for _, s := range solved {
        attempts, ok := attempts[s]
        if (ok) {
            score += 20 * attempts
        }
    }
    fmt.Print(len(solved))
    fmt.Print(" ")
    fmt.Println(score)
}
