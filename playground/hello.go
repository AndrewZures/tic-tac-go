package main

import (
  //      "os"
        "fmt"
        "bufio"
        "io"
       "bytes"
     )


func main() {
  var reader io.Reader
  var buffer bytes.Buffer

  //reader = os.Stdin
  reader = &buffer
  fmt.Fprintf(&buffer, "Hello\n1\n2\n")

  bio := bufio.NewReader(reader)
  for i := 0; i < 10; i++ {
  line, _, _ := bio.ReadLine()
  fmt.Println(string(line))
}

}
