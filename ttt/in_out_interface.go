package ttt


type InOutInterface interface {
  Read() string
  Println(string)
  Printf(string, ...interface{})
}


