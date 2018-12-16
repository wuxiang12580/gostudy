package main

import (
	"fmt"
	"readfile/file"
)
const filename = "readfile/abc.txt"

func main()  {
	fmt.Println(file.ReadFile(filename))
}
