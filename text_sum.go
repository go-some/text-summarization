package text_summarization

import (
  "github.com/JesusIslam/tldr"
)

type TextSum struct {
  extractor *tldr.Bag
}

func (self *TextSum) Init() {
  self.extractor = tldr.New()
}

func (self *TextSum) ExtractMainSents(text string, num int) []string {
  sentences, _ := self.extractor.Summarize(text, num)
  return sentences
}