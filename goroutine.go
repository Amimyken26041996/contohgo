package main

	

import "fmt"

	

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

	

func main() {
f("direct")

go f("goroutine")

go func(msg string){
fmt.println(msg)
} ("going")
var input string
fmt.Scanln(&input)
fmt.println("done")

}