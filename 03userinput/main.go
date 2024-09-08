package main 

import (
	"fmt" 
    "bufio"
    "os"
)

func main(){
	welcome := "Taking user inputs"
	fmt.Println(welcome);


	//READER
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating out of 5: \n")


	//comma,ok || err,err -> as golang does not have try catch so it expects us to handle it by variable (_ or err //choice)

	input, _ := reader.ReadString('\n');
	fmt.Println("Thanks for the rating ",input);
	

}