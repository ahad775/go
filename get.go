package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	
)

func main(){


	url:="http://localhost:3000/users"
	res,err:=http.Get(url)
	if err!=nil{
		panic(err)
	}

	defer res.Body.Close()

   response,err:=ioutil.ReadAll(res.Body)

   if err!=nil{
	   panic(err)
   }
   fmt.Println(string(response))
}