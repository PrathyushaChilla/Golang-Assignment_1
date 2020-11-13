package main 
   
import ( 
    "fmt"
    "encoding/json"
) 
   
type Human struct{ 
    Name string 
    Address string 
    Age int
} 
   
func main() { 
    var human1 Human 
    Data := []byte(`{ 
        "Name": "Deeksha",   
        "Address": "Hyderabad", 
        "Age": 21 
    }`) 
        err := json.Unmarshal(Data, &human1) 
       
    if err != nil { 
            fmt.Println(err) 
    } 
    fmt.Println("Struct is:", human1) 
    fmt.Printf("%s lives in %s.\n", human1.Name, human1.Address) 
       var human2 []Human    
    Data2 := []byte(` 
    [ 
        {"Name": "Vani", "Address": "Delhi", "Age": 21}, 
        {"Name": "Rashi", "Address": "Noida", "Age": 24}, 
        {"Name": "Rohit", "Address": "Pune", "Age": 25} 
    ]`) 
        err2 := json.Unmarshal(Data2, &human2) 
       
        if err2 != nil { 
                   fmt.Println(err2) 
        } 
       
        for i := range human2{ 
 fmt.Println(human2[i]) 
    } 
} 
output:
Struct is: {Deeksha Hyderabad 21}
Deeksha lives in Hyderabad.
{Vani Delhi 21}
{Rashi Noida 24}
{Rohit Pune 25}

Program exited.