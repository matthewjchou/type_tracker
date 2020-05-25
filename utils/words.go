package utils

import (
	"bufio"
	"os"
	"log"
	"math/rand"
	"time"
)

var words *[]string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getAllWords() *[]string {
	f, err := os.Open("resources/commonWords.txt")
	check(err)
	defer func () {
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}() // defer requires a function call, this defines an anonymous function and calls it
	words := make([]string, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text()) //append is funky, remember this!
	}
	return &words
}

func randomShuffle(N uint) *[]string {
	var words_set = make([]string, len(*words))
	copy(words_set, *words)
	rand.Seed(time.Now().UnixNano())
	//rand.Shuffle takes a swap function as the second parameter
	rand.Shuffle(len(words_set), func(i int, j int) {words_set[i], words_set[j] = words_set[j], words_set[i]})
	words_set = words_set[:N]
	return &words_set
}

func GetNWords(N uint) *[]string {
	if words == nil {
		//if first time running, get all the words
		words = getAllWords()
	}
	w := randomShuffle(N)

	return w
}