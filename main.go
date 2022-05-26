package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Function for word count, returns map
func word(s string) map[string]int {
	count := map[string]int{}
	for _, word := range strings.Fields(s)//Fields for words as a slice {
		count[word]++
	}
	return count
}

type count_word struct {
	word  string
	count int
}

func main() {
	// for input
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal()
	}

	result := word(str) //calling word count function  and assigning in result

	word_count := make([]count_word, 0, len(result)) // slice
	for k, v := range result {
		word_count = append(word_count, count_word{word: k, count: v}) //counting frequency of words
	}

	//sorting (high to low frequency of words)
	sort.Slice(word_count, func(i, j int) bool {
		return word_count[i].count > word_count[j].count
	})

	for i := 0; i < 10 && i != len(word_count); i++ {
		fmt.Println("Word ", word_count[i].word, "\t count :  \t", word_count[i].count)

	}

}
