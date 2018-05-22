package main
import "fmt"

var g = make(chan int)
var quit = make(chan chan bool)
//为了等待协程B退出 所以把wait这个chan压入quit中
func main() {
   go B()
   for i := 0; i < 5; i++ {
       g <- i
       fmt.Println("i == ", i)
   }
   wait := make(chan bool)
   quit <- wait
   value := <-wait //这样就可以等待B的退出了
   fmt.Println("value = ", value)
   fmt.Println("Main Quit")
}

func B() {
   for {
       select {
       case i := <-g:
           fmt.Println(i + 1)
       case c := <-quit:
           c <- true
           fmt.Println("B Quit")
           return
       }
   }
}
