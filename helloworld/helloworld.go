package main

import "fmt"

func main() {
  
  // function test
  foo()

  // sequence test
  fmt.Println("My first Hello World usig GoLang!")

  // loop test
  for i := 0; i <= 10; i++ {
    fmt.Printf("loop test i = %d\n", i)
    // condition test
    if i%2 == 0  { 
      fmt.Printf("condition test i = %d\n", i)
    }
  }

  // function test
  bar()
}

func foo() {
  fmt.Println("Enter main()")
}

func bar() {
  fmt.Println("Exit main()")
}
