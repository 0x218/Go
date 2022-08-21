//go build --help
//This serve as the library for C++.  This file need to build first with below comamnd -
//go build -buildmode=c-shared -o ./golibServer.so ./golibServer.go

package main

import (
	"C"
	"fmt"
	"io/ioutil"
)

//export GoConcatenate
func GoConcatenate(sIn string, bIn []byte, bOut []byte) {
	n := copy(bOut, sIn)
	copy(bOut[n:], bIn) //saving it to the out variable.
}

func joinStrings(systemInput string, userInput string, forwardJoin bool) string {
	if forwardJoin == true {
		return systemInput + userInput
	} else {
		return userInput + systemInput
	}
}

//export GoSayHello
func GoSayHello(cinput *C.char) *C.char {
	userInput := C.GoString(cinput)
	return C.CString(joinStrings("Hello ", userInput, true)) //add Hello before the userinput
}

//export GoPrintFileContent
func GoPrintFileContent(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File read err")
	}
	fmt.Println(string(content))
}

func main() {

}
