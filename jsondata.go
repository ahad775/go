package main

import (
	"encoding/json"
	"fmt"
)


type  Student struct{
Name string 
Age int
RollNo int 
FatherName string
Subjects []string
}
func main(){

user:=[]Student{
	{
		Name: "John Doe",
        Age: 20,
        RollNo: 1,
        FatherName: "Jane Doe",
        Subjects: []string{"Math", "Science", "English"},
	},{
		Name: "Kon moe",
        Age: 20,
        RollNo: 1,
        FatherName: "Kim jon",
        Subjects: []string{"Math", "Science", "English","dictatorshipism"},
	},
}

json,err:=json.Marshal(user)

if err!=nil{
	panic(err)
}
fmt.Println(string(json))

}

