package main

import (
	"fmt"
  term "github.com/nsf/termbox-go"
  "github.com/goml/gobrain"
  "math/rand"
)

func reset (){
  term.Sync()
  fmt.Println("Select :  Esc or q exit.")
  fmt.Println("1. Hello World.")
  fmt.Println("2. Triangle.")
  fmt.Println("3. gobrain example.")
  fmt.Println("4. func scope example.")
}


func p1(){
  fmt.Println("Hello, World!")
}

func p2(s int) {
  for i := 0 ; i < s ; i++ {
    for j := 0 ; j < s - i ; j++ {
      fmt.Print(" ")
    }
    for k := 0 ; k < ( i * 2 ) + 1 ; k++ {
      fmt.Print("*")
    } 
    fmt.Println()
  } 
}

func p3(){
  rand.Seed(0)
  patterns := [][][] float64 {
            {{0,0},{0}},
            {{0,1},{1}},
            {{1,0},{1}},
            {{1,1},{0}},
  }
  ff := &gobrain.FeedForward{}
  ff.Init(2,2,1)

  ff.Train(patterns,1000,0.6,0.4,true)

  ff.Test(patterns)

  inputs := []float64{1,1}
  fmt.Println("{1,1}=>", ff.Update(inputs))
}

func p4(){
  add := func(a,b int) int {
    return a+b;
  }
  fmt.Println("add(3,4)=", add(3,4))
  foo := scope()
  fmt.Println("scope:", foo())
  varOne  , varTwo := outer()
  fmt.Println("outer:", varOne(),varTwo)
}

func scope() func() int {
  outerVar := 42
  foo := func() int { return outerVar }
  return foo
}

func outer () (func() int , int) {
  outerVar := 2
  inner := func() int {
    outerVar += 99
    return outerVar
  }
  return inner, outerVar
}

func main() {
  err := term.Init()
  if err != nil {
    panic(err)
  }
  defer term.Close()
  reset()
  keyPressLoop:
  for {
    switch ev:= term.PollEvent(); ev.Type {
      case term.EventKey:
        switch ev.Key {
          case term.KeyEsc:
            break keyPressLoop
          default:
            reset()
            fmt.Println("ASCII: " , ev.Ch)
            switch ev.Ch {
            case '1':
              p1()
            case '2':
              p2(5)
            case '3':
              p3()
            case '4':
              p4()
            case 'q':
              break keyPressLoop
            }
        }
    }
  }
}
