package main

import (
    "fmt"
    "strings"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    counts := make(map[string]int)

    for _, word := range words {
        counts[word]++
    }

    return counts
}

func main() {
    fmt.Println(WordCount("I am learning Go!"))
    fmt.Println(WordCount("The quick brown fox jumped over the lazy dog."))
    fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
    fmt.Println(WordCount("A man a plan a canal panama."))
}
