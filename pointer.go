package main
import "fmt"
func zeroval(ival int){
   ival = 0
}

func Zeroval (iptr * int) {
     * iptr = 0
}

func main () {
    i : 1 
     fmt.println("initial :",i)

Zeroptr (i)
fmt.println("zeroval :",i)

Zeroptr (&i)
fmt.println("zeroptr:",i)

fmt.prntl(" pointer :",&i)


}