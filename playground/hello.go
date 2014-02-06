package main

import (
      //  "os"
        "fmt"
        "bufio"
        "io"
        "bytes"
      )


func main() {
  var reader io.Reader
  var buffer bytes.Buffer

  reader = &buffer
  fmt.Fprintf(&buffer, "Hello")

  //buf := new(bytes.Buffer)
  bio := bufio.NewReader(reader)
  //buf.ReadFrom(os.Stdin)
  line, _, _ := bio.ReadLine()
  fmt.Println(string(line))

}
