program1
package main
import (  
    "fmt"
    "os"
)

func main() {  
    f, err := os.Create("test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString("Hello World")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}
output:
11 bytes written successfully

Program exited.


program2
package main

import (  
    "fmt"
    "os"
)

func main() {  
    f, err := os.Create("/home/prathyusha/bytes")
    if err != nil {
        fmt.Println(err)
        return
    }
    d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
    n2, err := f.Write(d2)
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(n2, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
return
    }
}
output:
open /home/prathyusha/bytes: no such file or directory

Program exited.




program3
package main

import (  
    "fmt"
    "os"
)

func main() {  
    f, err := os.Create("lines")
    if err != nil {
        fmt.Println(err)
                f.Close()
        return
    }
    d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}

    for _, v := range d {
        fmt.Fprintln(f, v)
        if err != nil {
            fmt.Println(err)
            return
        }
    }
    err = f.Close()
	if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("file written successfully")
}
output:
file written successfully

Program exited.