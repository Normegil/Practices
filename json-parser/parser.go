package main

import (
  "encoding/json"
  "fmt"
)

type SubTest struct {
  Test1 string
}

type test struct {
  Test []SubTest
}

func main() {
  content := []byte(`{"test": [{"Test1" : "XY"}, {"Test1" : "ZA"}]}`)
  var testVar test
  err := json.Unmarshal(content, &testVar)
  if nil != err {
    panic(err)
  }
  fmt.Println(testVar.Test)
}
