package main

import (
	"fmt"
	s "strings"
	"log"
)

func SpinWords(str string) string {
	if s.Contains(str, " ") {
  var splited []string = s.Fields(str)
  log.Println(splited[0]);
  
  for i := 0; i <= len(splited)-1; i++ {
    if len(splited[i]) >= 5 {
      //var wordSplitInLetters = s.Split(splited[i], "")
      var wordSplitInLetters = []rune(splited[i])
      
      for j,k := 0, len(wordSplitInLetters)-1; j<k; j,k =j+1, k-1 {
        wordSplitInLetters[j], wordSplitInLetters[k] = wordSplitInLetters[k], wordSplitInLetters[j]
      }
      var wordSpined = string(wordSplitInLetters)
      splited[i] = wordSpined
    }
  }
  return s.Join(splited, " ")
  } else {
    if len(str) >= 5 {
      //var wordSplitInLetters = s.Split(splited[i], "")
      var wordSplitInLetters = []rune(str)
      
      for j,k := 0, len(wordSplitInLetters)-1; j<k; j,k =j+1, k-1 {
        wordSplitInLetters[j], wordSplitInLetters[k] = wordSplitInLetters[k], wordSplitInLetters[j]
      }
      var wordSpined = string(wordSplitInLetters)
 str = wordSpined
}
}// SpinWords
return str
}

func main() {

	fmt.Println(SpinWords("Test Sample"));
	fmt.Println(SpinWords("Sample"));
	fmt.Println(SpinWords("Test Sample More Test Sample"));
}


//j=0, 5-1-0=4
//j=1, 5-1-1=3
//j=2, 5-1-2=2
//j=3, 5-1-3=1
//j=4, 5-1-4=0
