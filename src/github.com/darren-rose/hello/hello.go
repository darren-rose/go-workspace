package main

import (
  "fmt"
  "math"
)

type Person struct {
  first, last string
}

// add function to Person that returns a string
func (p Person) greet() string {
  return "Hello, I am " + p.first + " " + p.last
}

// add function to Person that updates Person last
func (p *Person) marries(last string) {
  p.last = last
}

type Shape interface {
  area() float64
}

type Circle struct {
  radius  float64
}

func (c Circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

type Rectangle struct {
  x, y  float64
}

func (r Rectangle) area() float64 {
  return r.x * r.y
}

func getArea(s Shape) float64 {
  return s.area()
}

func adder() func(int) int {
  sum := 0;
  return func(x int) int {
    sum += x
    return sum
  }
}

func main() {

  c := Circle{5}
  r := Rectangle{5, 10}
  fmt.Println(getArea(c))
  fmt.Println(getArea(r))

  name, size := "Ann", 1.1

  add  := adder()

  fmt.Println(add(1))
  fmt.Println(add(9))

  p1 := Person{"Mary", "Rose"}

  fmt.Println(p1)
  fmt.Println(p1.greet())
  p1.marries("Jones")
  fmt.Println(p1.greet())

  a := 5
  b := &a

  *b = 7

  fmt.Println(name, size, math.Floor(2.7), math.Ceil(2.7), math.Sqrt(16))
  fmt.Printf("%T %T\n",  name, size)
  fmt.Println(a, b)
  fmt.Println(a, *b)
  fmt.Printf("%T %T\n",  a, b)

}
