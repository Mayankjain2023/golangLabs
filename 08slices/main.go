package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Learning slices")

	var fruitList = []string{} //when using this syntax need to initialize it
	//we can add as many values as we like and it expands automatically for us -> similar to arrays in other lang
	fmt.Printf("Type of the fruitlist %T\n", fruitList)

	//in case of array you add values by accessing index postion
	//in slices you use append(slice you want to modify, value to be added)

	fruitList = append(fruitList, "Mango")
	fruitList = append(fruitList, "Grapes", "Berries", "Apple")

	fruitList = append(fruitList[:3])
	// slicing the slice -> modifies the list to include elements as per specified index
	//range is non-inclusive so 1,3 will include elements 1,2 only

	fmt.Println("Modified list", fruitList)

	highScore := make([]int, 5) // creating slices using make , need to specify the type and size
	highScore[0] = 1
	highScore[1] = 123
	highScore[2] = 345
	highScore[3] = 0
	highScore[4] = 9999
	// highScore[5] = 99995; you can't do this bcz you declared size 4
	//but you can do this shit ...
	highScore = append(highScore,55555,67,234,123456,76,2)
	fmt.Println("Highscores", highScore)
	sort.Ints(highScore)
	fmt.Println("Highscores sorted", highScore)
	fmt.Println("Are scores sorted?. ",sort.IntsAreSorted(highScore))

	//how to remove a value from slices via index

	 courses := make([]string,0)
	 courses = append(courses,"c++","c","reactJs","nodejs","golang");

	 fmt.Println("Complete list of courses",courses);
	 var index int = 2;
	 courses = append(courses[:index],courses[index+1:]... )
	 fmt.Println("Modified list of courses",courses)




}
