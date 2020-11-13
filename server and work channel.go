program1 Working with channels

package main
 import (
    "fmt"
)
 func SendDataToChannel(ch chan int, value int) {
    ch <- value
}
 func main() {
     var v int
    ch := make(chan int)     
     go SendDataToChannel(ch, 101)   
     v = <-ch                         
    fmt.Println(v)       
}
output:
101





program2 Sending custom data via channels

Program exited.
mport (
    "fmt"
    // "time"
)
type Person struct {
    Name string
    Age  int
}
func SendPerson(ch chan Person, p Person) {
    ch <- p
}
func main() {
     p := Person{"prathyusha", 23}
     ch := make(chan Person)
     go SendPerson(ch, p)
     name := (<-ch).Name
    fmt.Println(name)
}
output:
prathyusha

Program exited.





program3 Closing a Channel

package main
 import "fmt"
 func SendDataToChannel(ch chan string, s string) {
    ch <- s
    close(ch)
}
 func main() {
     ch := make(chan string)
     go SendDataToChannel(ch, "Hello World!")
 v, ok := <-ch
     if !ok {
        fmt.Println("Channel closed")
    }
     fmt.Println(v) 
	output:
	Hello World!

Program exited.