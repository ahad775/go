

package main
import "fmt"
func main(){
/* var b int=add(23,45) */
c:=clouser()
fmt.Println(c(1))
fmt.Println(c(1))
fmt.Println(c(1))
fmt.Println(c(1))
fmt.Println(c(1))
fmt.Println(c(1))


mult:=func(a int , b int) int{
	return a*b
}
var a int=mult(23,45)
fmt.Println(a)
}

func add(a int ,b int) int {
	return a+b
}

func clouser() func(int) int{
	counterMul:=1
	return func(x int)int {
	counterMul=x+counterMul
    return  counterMul
	}
}
