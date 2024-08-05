package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	
)

func main() {
	res, err := http.Get("http://localhost:3000/users")
	if err!=nil{
		return
	}
	defer res.Body.Close()
	txt,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		return
	}
	fmt.Println(string(txt))
}