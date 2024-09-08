package main

import("fmt")

func main(){
	fmt.Println("Arrays ")


	var fruitList[4] string;

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"

	fmt.Println("Fruitlist",fruitList)
	fmt.Println("Length of the fruitlist",len(fruitList))


	var vegList = [3]string{"Potato","Beans","Tomato"}

	fmt.Println("Veg list length ",len(vegList))
}