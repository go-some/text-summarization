package main

import (
	"fmt"
	"io/ioutil"
	"github.com/JesusIslam/tldr"
)

func main() {
	intoSentences := 3
	textB, _ := ioutil.ReadFile("./input/sample.txt")
	text := string(textB)
	bag := tldr.New()
	result, _ := bag.Summarize(text, intoSentences)
	fmt.Println(result)
}
