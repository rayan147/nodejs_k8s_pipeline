package main

import (
   "fmt"
   "reflect"
)
 const (
      Pi = 3.14
      Avogadro float32 = 6.022e23
      
   )
func main()  {

   var str string = "Hello World"
   var b bool = true
   var i int = 10
   var f float32 = 1.23
   var c complex64 = 1 + 2i
   var a [5]int = [5]int{1,2,3,4,5}
   var s []int = []int{1,2,3,4,5}
   var m map[string]int = map[string]int{"one":1, "two":2}
   var p *int = &i
   var r rune = 'a'
   var n uint = 10
   var ui uint = 10
   var ui8 uint8 = 10
   var ui16 uint16 = 10
   var ui32 uint32 = 10
   var ui64 uint64 = 10
   var i8 int8 = 10
   var i16 int16 = 10

   fmt.Println(str)
   fmt.Println(b)
   fmt.Println(i)
   fmt.Println(f)
   fmt.Println(c)
   fmt.Println(a)
   fmt.Println(s)
   fmt.Println(m)
   fmt.Println(p)
   fmt.Println(r)
   fmt.Println(n)
   fmt.Println(ui)
   fmt.Println(ui8)
   fmt.Println(ui16)
   fmt.Println(ui32)
   fmt.Println(ui64)
   fmt.Println(i8)
   fmt.Println(i16)





   fmt.Println("Hello, World!")
   fmt.Println(Pi)
   fmt.Println(Avogadro)
   fmt.Println(reflect.TypeOf(Pi))
   fmt.Println(reflect.TypeOf(Avogadro))
}