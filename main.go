package main

import "fmt"

// This is a comment
/*
This is my own guide to learn Go
*/

func main()  {
    fmt.Println("Hello World")
    exercice10()

}

func other(){
  fmt.Println("Helloooo!")
  fmt.Println("1+1=",1+1)
  fmt.Println(len("Steven"))
  fmt.Println("Steven"[0])
}

func exercice1(){

  // var nombre string  = "Steven"
  // edad := 15
  //fmt.Println(nombre + " tiene "+ edad + "años")
}

func exercice2(){
  x := 1
  for x <= 10 {
    fmt.Println(x)
    x += 1
  }
  for i:= 0;i <= 10; i++ {
      fmt.Println(i)
      if i% 2 == 0 {
        fmt.Println("is even")
      }
  }

}


func exercice3(){
    var numbers [4]int
    numbers[0] = 1
    fmt.Println(numbers)

}

func print(a string){
  fmt.Println(a)
}

func exercice4(){
  // var ss []int
  // This actually fails var test map[string]int
  test := make(map[string]int)
  test2 := map[string]string{}
  test["key"] = 10
  test["age"] = 50
  test2["steven"] = "sierra"
  fmt.Println(test)
  fmt.Println(test2)
  item,ok := test2["steven"]
  fmt.Println(item,ok)
  print(item)

}

func exercice5(){
  panic("Boom!")
}

func exercice6(){
  x := 5
  zero(&x)
  fmt.Println(x)

}

func exercice7(){

  fmt.Println(multipleValues(1))
}

func  exercice8(){
  f1 := func(a int, b int) int{
    return a+b
  }
  r := f1(1,3)
  fmt.Println("El resultado es",r)
}

func exercice9() []int {
  return testReturn()
}

func exercice10(){
   v := testReturn()
   for _,item := range v {
     fmt.Println(item)
   }
}

func testReturn() []int {
  shit := []int{1,2,3}
  return shit

}

func multipleValues(arg int) (int,string)  {

 return arg,"jajajja!"
}

func zero(numero *int){
    *numero = 0
}