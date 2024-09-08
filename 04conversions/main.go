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

	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err!=nil {
		fmt.Println("Error occured",err);
	}else{
		fmt.Println("Thanks for the rating ",numRating + 1);
		fmt.Printf("Type of rating %T",numRating);
	}

}