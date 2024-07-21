// Input: aaabbcaac // Output: 3a 2b 1c 2a 1c -> pairs {counter, symbol}

package main

import (
  "fmt"
  "strconv"
)

func main() {
  str := "aaabbcaac"
  cnt := 0
  var res []string
  for i := 0; i < len(str)-1; i++ {
    cnt++
    if str[i] != str[i+1] {
      if i == len(str)-2 {
        res = append(res, strconv.Itoa(cnt)+string(str[i]))
        res = append(res, strconv.Itoa(1)+string(str[i+1]))
        break
      }
      res = append(res, strconv.Itoa(cnt)+string(str[i]))
      cnt = 0
    }
  }
  fmt.Println(res)
}