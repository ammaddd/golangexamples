package golangexamples

import (
	"fmt"
	//"greetings"
	"github.com/ehteshamz/greetings"
)

func ConcatSlice(sliceToConcat []byte)string {
    s:=""
    for i:= range sliceToConcat {
       	s=s+"-"+string(sliceToConcat[i])
        }
    return s
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int){
    for i:= range sliceToEncrypt {
	sliceToEncrypt[i]=sliceToEncrypt[i]+byte(ceaserCount)
        }
    fmt.Println(string(sliceToEncrypt))

}

func EZGreetings(name string) string{
    return greetings.PrintGreetings(name)

}

