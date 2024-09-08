package main 


import ("fmt")


func main (){
	fmt.Println("Playing with pointers ")

	// var ptr *int 
	// fmt.Println("Valule of pointer is",ptr)
	// fmt.Printf("type of ptr is%T",ptr)

	myNumber := 24
	var ptr = &myNumber // when referencing to a memory location & is used
	fmt.Println("Address of mynumber is ",ptr);
	fmt.Println("Address of mynumber is ",*ptr);

	*ptr = *ptr * 2;

	fmt.Println("New value of mynumber is ",myNumber);

}




//Memory allocation and deallocation in golang

/*

1) Memory allocation and deallocation happens automatically in golang
2) new () and make() -> ways to allocate memory
3) new () : -> not initialize
	> Allocate memory but no INIT
	> you will get a memory address
	> zeroed storage -> cannot put any data initially
4) make () :
	>Allocate memory and INIT
	>you will get a memory address
	>non-zeroed storage -> can put any data initially
5) make() is recommended and is used most of the times
6) Garbage collection happens automatically -> after a certain threshold 


*/