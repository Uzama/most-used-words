package processor

import "strings"

// custom type to keep track of a word and it occurrence
type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

// process the text and return 10 most used words
func TextProcessor(text string) []WordCount {

	c := make(chan string)

	go generateWords(text, c)

	wordCount := make(map[string]WordCount)

	// recive words from the channel and counts the no of occurrence
	for word := range c {

		lower := strings.ToLower(word)

		val, ok := wordCount[lower]
		if !ok {
			wordCount[lower] = WordCount{Word: lower, Count: 1}
			continue
		}

		val.Count++
		wordCount[lower] = val
	}

	result := make([]WordCount, 0)

	// build array to do heap sort
	for _, wc := range wordCount {
		result = append(result, wc)
	}

	return HeapSort(result)
}

// generate words from given text and share the words via channel
func generateWords(text string, c chan string) {

	flag := false // keep track whether word started
	word := ""    // keep track of the characters of a word

	for _, char := range text {

		// check whether the character is alphabats
		isAlphabat := ('a' <= char && 'z' >= char) || ('A' <= char && 'Z' >= char)

		if !flag && !isAlphabat {
			continue
		}

		// start of the word
		if !flag && isAlphabat {
			word += string(char)
			flag = true
			continue
		}

		// keep concatenating characters to build word
		if flag && isAlphabat {
			word += string(char)
			continue
		}

		// sending the word via channel
		c <- word

		// set to defualt before starting with next word
		word = ""
		flag = false
	}

	// if word is not empty send the word via channel (last word in the text)
	if word != "" {
		c <- word
	}

	// close the channel to notify finished sending all words
	close(c)
}
