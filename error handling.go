program1

package main

import (
    "errors"
    "fmt"
)

func main() {
    sampleErr := errors.New("error occured")
    fmt.Println(sampleErr)
}
output:
error occured
Program exited.





program2
package main

import "errors"
import "fmt"
import "math"

func MathFunction (value float64)(float64, error) {
   if(value < 0){
      return 0, errors.New("Math: negative number passed to Sqrt")
   }
   return math.Sqrt(value), nil
}

func main() {
   result, err:= MathFunction (-1)

   if err != nil {
      fmt.Println(err)
   }else {
      fmt.Println(result)
   }
   
   result, err = Sqrt(9)

   if err != nil {
      fmt.Println(err)
   }else {
      fmt.Println(result)
   }
}
