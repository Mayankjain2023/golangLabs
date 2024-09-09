package main

import "fmt"

func main(){

	//we dont have classes in golang 
	//we dont have inheritance , super, parent etc here as they are concepts of the classes
	fmt.Println("Getting started with structs")

	mayank := User{"Mayank","mayank@go.dev",true,23}
	fmt.Println("User Mayank",mayank);
	fmt.Printf("Mayank details %+v\n",mayank);
	fmt.Printf("Name is %+v\n",mayank.Name);


}


//struct

type User struct{
	Name string
	Email string
	Status bool
	Age int
}