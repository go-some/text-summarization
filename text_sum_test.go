package text_summarization

import (
  "testing"
  "io/ioutil"
)

func TestExtractMainSents(t *testing.T) {
  
  textSum := TextSum{}
  textSum.Init()
  
  num := 3
  text, _ := ioutil.ReadFile("./input/sample.txt")
  textList := string(text)
  sentences := textSum.ExtractMainSents(textList, num)
  
  if len(sentences) != num {
    t.Errorf("[length error] \n num is %d but length of extracted sentences is %d",
              num,
              len(sentences),
            )
  }
}
