package main

	

import "fmt"
import "os"

func main(){


f := createFile("/tmp/defer.txt")

defer closeFile(f)
writeFile(F)

}

func createFile(p string) *os.file{
fmt.println("creating")
f, err := os.Create(p)
if err != nil{
panic(err)
}

return f

}

func writeFile(f * os.file) {
fmt.println("writing")
fmt.Fprintln("closing")
f.close()

}