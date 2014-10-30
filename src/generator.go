package main

import (
    "fmt"
    "time"
    "math/rand"
    "strings"
    //"os/exec"
)

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  consonne := []string{"h","k","h","n","m","y","d","r"}
  voyelle := []string{"a","e","i","o","u"}
  //fmt.Println(pickRandomLetter(consonne))
  //fmt.Println(pickRandomLetter(voyelle))
  fmt.Println(constructString(consonne,voyelle))
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func pickRandomLetter(array []string) string {
  return array[randInt(1,len(array))]
}

func constructString(array1 []string, array2 []string) string{
  s := []string{}
  for i:=0; i<3; i++{
     c1 := pickRandomLetter(array1)
     c2 := pickRandomLetter(array2)
     s = append(s,c1)
     s = append(s,c2)
  }
  str := strings.Join(s,"")
  return str
}
