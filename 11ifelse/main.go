package main

import "fmt"

func main(){

	fmt.Println("If else shit")


	loginCount := 23
	var message string

	if loginCount<10 {
		message = "Less count"
	}else if loginCount == 15 {
		message = "Mediocrity"
	}else if loginCount == 20{
		message = "decent"
	}else {
		message = "gOOD"
	}

	fmt.Println("result",message);



	if 9%2 == 0{
		fmt.Println("Number is even")
	}else{
		fmt.Println("Number is odd")
	}

	if num := 32;num<40{
		fmt.Println("Number is lesser than 40")
	}else{
		fmt.Println("Number is greater than 40")
	}
}