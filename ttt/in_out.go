package ttt

type InOut interface {
  Read() string
  Println(string)
  Printf(string, ...interface{})
}


