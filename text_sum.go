package main

import (
  "fmt"
  "time"
  "log"
  "io/ioutil"
  "github.com/JesusIslam/tldr"
)

type TextSum struct {
  extractor *Bag
}

func (self *TextSum) init() {
  self.extractor = tldr.New()
}

func (self *TextSum) extract_main_sents(sentences string, num int) string {
  result, _ := self.extractor.Summarize(sentences, num)
  return result
}