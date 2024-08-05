package main
import (
	"fmt"
)
func main(){
slc:=[]int{1,2,3,4}
for i,v:=range slc{
	fmt.Println(i,v)
}
s:=slc[1:3]
b:=append(s,4)
fmt.Println(s)
fmt.Println(b)
}