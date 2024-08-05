package main 


import (
	"fmt"
)
func main(){

mymap:=make(map[string]string)

mymap["name"]="ahad"

fmt.Println(mymap["name"])
for key,value:=range mymap{
	fmt.Println(key,value)
}
}