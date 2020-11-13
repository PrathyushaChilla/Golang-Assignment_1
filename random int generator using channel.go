1.package main 
  
import ( 
    "fmt"
    "math/rand"
    "time"
) 
  
func main() { 
  
    fmt.Println(rand.Intn(200)) 
    fmt.Println(rand.Intn(200)) 
    fmt.Println(rand.Intn(200)) 
    fmt.Println() 
    fmt.Println(rand.Float64()) 
  
    
    fmt.Println((rand.Float64() * 9) + 10) 
    fmt.Println((rand.Float64() * 8) + 9) 
    fmt.Println() 
    s1 := rand.NewSource(time.Now().UnixNano()) 
    p1 := rand.New(s1) 
      
    fmt.Println(p1.Intn(200)) 
    fmt.Println(p1.Intn(200)) 
    fmt.Println() 
  
    s2 := rand.NewSource(55) 
    p2 := rand.New(s2) 
    fmt.Println(p2.Intn(200)) 
    fmt.Println(p2.Intn(200)) 
    fmt.Println() 
      
    s3 := rand.NewSource(5) 
    p3 := rand.New(s3) 
    fmt.Println(p3.Intn(200)) 
    fmt.Println(p3.Intn(200)) 
} 
$go run main.go
81
87
47

0.4377141871869802
13.821737473641392
14.494584582936875

175
177

112
164

26
36






2.package main 
  
import ( 
    "crypto/rand"
    "fmt"
) 
  
func main() { 
    RandomCrypto, _ := rand.Prime(rand.Reader, 134) 
    fmt.Println(RandomCrypto) 
} 
$go run main.go
21217911299285046682359992173637710442797
© 2020 GitHub, Inc.
Terms
Privacy
Security
Status
Help
Contact GitHub
Pricing
API
Training
Blog
About
