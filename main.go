package main

import (
  "fmt"
  "time"
  "log"
  "io/ioutil"
  "github.com/JesusIslam/tldr"
)

func main() {
  intoSentences := 3

  start := time.Now()
  
  textB, _ := ioutil.ReadFile("./input/sample.txt")
  text := string(textB)
  bag := tldr.New()
  result, _ := bag.Summarize(text, intoSentences)
  fmt.Println(result)
  
  elapsed := time.Since(start)
  log.Printf("It took %s", elapsed)
}
