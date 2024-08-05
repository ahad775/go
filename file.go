package main

import "os"

func main() {
	file,err := os.Create("./t.txt")
	if err!=nil{
	return
	}
	defer file.Close()

	_,err=file.WriteString("hey")
	if err!=nil{
		return
		
	}
}