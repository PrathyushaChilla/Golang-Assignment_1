package main

    import (
        "fmt"
        "strings"
        "unicode"
    )

    func verifyPassword(password string) error {
        var uppercasePresent bool
        var lowercasePresent bool
        var numberPresent bool
        var specialCharPresent bool
        const minPassLength = 8
        const maxPassLength = 64
        var passLen int
        var errorString string

        for _, ch := range password {
            switch {
            case unicode.IsNumber(ch):
                numberPresent = true
                passLen++
            case unicode.IsUpper(ch):
                uppercasePresent = true
                passLen++
            case unicode.IsLower(ch):
                lowercasePresent = true
                passLen++
case unicode.IsLower(ch):
                lowercasePresent = true
                passLen++
            case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
                specialCharPresent = true
                passLen++
            case ch == ' ':
                passLen++
            }
        }
        appendError := func(err string) {
            if len(strings.TrimSpace(errorString)) != 0 {
                errorString += ", " + err
            } else {
                errorString = err
}
        }
        if !lowercasePresent {
            appendError("lowercase letter missing")
        }
        if !uppercasePresent {
            appendError("uppercase letter missing")
        }
        if !numberPresent {
            appendError("atleast one numeric character required")
        }
        if !specialCharPresent {
            appendError("special character missing")
        }
        if !(minPassLength <= passLen && passLen <= maxPassLength) {
            appendError(fmt.Sprintf("password length must be between %d to %d characters long", minPassLength, maxPassLength))
        }

        if len(errorString) != 0 {
            return fmt.Errorf(errorString)
        }
        return nil
    }
 func main() {
        password := "Prathyusha"
        err := verifyPassword(password)
        fmt.Println(password, " ", err)
    }
	output:
	Prathyusha   atleast one numeric character required, special character missing

Program exited.



program2
ackage main 
import (  
    "fmt"  
    "regexp"  
)  
  
func main() {  
    re := regexp.MustCompile("f([a-z]+)ing")  
    fmt.Println(re.FindStringSubmatch("fprathyushaing"))  
    fmt.Println(re.FindStringSubmatch("abcfloatingxyz"))  
}  
output:
[fprathyushaing prathyusha]
[floating loat]

Program exited.