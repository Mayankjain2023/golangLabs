package main

import ("fmt")

// no var syntax 
// numberBig := 3000000.33; //not allowed outside the method


//const
const LogginToken string = "ghibrish12345"  //Capital L signifies thats its a public variable


func main(){
	//string
	var username string = "Mayank"
	fmt.Println(username);
	fmt.Printf("Variable is of type : %T \n",username);

	//boolean
	var isLoggedIn bool 
	fmt.Println(isLoggedIn);
	fmt.Printf("Variable is of type : %T \n",isLoggedIn);

	//small-int
	var smallVal uint8 = 255
	fmt.Printf("Variable is of type : %T \n",smallVal);

	//small-float
	var smallFloat float32 = 255.554433221199887766
	fmt.Println(smallFloat);
	fmt.Printf("Variable is of type : %T \n",smallFloat);

	//int
	//default value is 0
	var number int
	fmt.Println(number);
	fmt.Printf("Variable is of type : %T \n",number);


	//why declaring type is important
	var website = "learning.com"
	fmt.Println(website);
	//website = 4  // it wont let you do this because it's already considered at string
	fmt.Printf("Variable is of type : %T \n",website);


	// no var syntax 
	numberBig := 3000000.33; 
	fmt.Println(numberBig);

	//loggin
	
}