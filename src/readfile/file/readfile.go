package file

import (
	"io"
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) string  {
	file,err :=os.Open(filename)
	if err!=nil{
		panic(err)
	}
	readFileContent(file)
	return "ok"
}
func readFileContent(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

