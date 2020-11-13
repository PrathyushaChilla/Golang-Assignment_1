program1 struct

package main

import (  
    "fmt"
)

type Address struct {  
    city  string
    state string
}

type Person struct {  
    name    string
    age     int
    address Address
}

func main() {  
    p := Person{
        name: "supraja",
        age:  26,
        address: Address{
            city:  "rajampet",
            state: "andhrapradesh",
        },
    }

    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.address.city)
    fmt.Println("State:", p.address.state)
}
output:
Name: supraja
Age: 26
City: rajampet
State: andhrapradesh
Program exited.