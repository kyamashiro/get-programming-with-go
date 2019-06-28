package main

import (
	"fmt"
	"strings"
)

var article = "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."

func main() {
	splitWords := strings.Fields(strings.ToLower(article))

	counter := make(map[string]int)

	for _, word := range splitWords {
		word = strings.Trim(word, ",.;—")
		counter[word]++
	}

	for word, num := range counter {
		if 2 <= num {
			fmt.Println(word, num)
		}
	}
}
