package main

import "fmt"

// Имя интерфейса и структуры совпадает
type MyInterface interface {
    SomeMethod()
}

type MyStruct struct {
    Name string
}

func (s MyStruct) SomeMethod() {
    fmt.Println("Hello from SomeMethod")
}

func main() {
    myStruct := MyStruct{Name: "Example"}
    myStruct.SomeMethod()
}
