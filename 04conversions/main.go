package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)


func main(){

	fmt.Println("Welcome to our pizza shop!")
	fmt.Println("Please do rate our pizza between  1-5");


	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')

	fmt.Println("Thanks for the rating ",input);
	//cannot directly assign /add number to string so need to convert it to compatible type
	//while passing the string \n was also going along with it so need to trim the space in input
	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)


	//handle the error properly
	if err!=nil {
		fmt.Println("Error occured",err);
	}else{
		fmt.Println("Thanks for the rating ",numRating + 1);
		fmt.Printf("Type of rating %T",numRating);
	}

}