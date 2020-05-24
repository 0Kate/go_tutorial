package main

import (
    "strings"
    "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
    wordCountMap := make(map[string]int)

    words := strings.Fields(s)
    for i := 0; i < len(words); i++ {
        wordCountMap[words[i]]++
    }

    return wordCountMap
}

func main() {
    wc.Test(WordCount)
}
