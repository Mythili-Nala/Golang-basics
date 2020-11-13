package main

import "fmt"

func main() {

   var a int = 50
   var b int = 30
   var c int

   c = a + b
   fmt.Printf("Addition - Value of c is %d\n", c )
   
   c = a - b
   fmt.Printf("Subtraction - Value of c is %d\n", c )
   
   c = a * b
   fmt.Printf("Multiplication - Value of c is %d\n", c )
   
   c = a / b
   fmt.Printf("Division - Value of c is %d\n", c )
   
   c = a % b
   fmt.Printf("Mod - Value of c is %d\n", c )
   
   a++
   fmt.Printf("Increment - Value of a is %d\n", a )
   
   a--
   fmt.Printf("Decrement - Value of a is %d\n", a )
}