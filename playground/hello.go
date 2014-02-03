package main

import "fmt"

type Shaper interface {
  Area() int
}

type Rectangle struct {
	length, width int
}

type Square struct {
	side int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (sq Square) Area() int {
	return sq.side * sq.side
}

func Add (x, y int) (sum int) {
    return x + y
}

func main() {
  r := Rectangle{length:5, width:3}
  fmt.Println("Rectangle details are: ",r)

  var s Shaper
  s = r //equivalent to "s = Shaper(r)" since Go identifies r matches the Shaper interface
  fmt.Println("Area of Shaper(Rectangle): ", s.Area())
  fmt.Println()

  q := Square{side:5}
  fmt.Println("Square details are: ",q)
  s = q //equivalent to "s = Shaper(q)
  fmt.Println("Area of Shaper(Square): ", s.Area())
  fmt.Println()

}
