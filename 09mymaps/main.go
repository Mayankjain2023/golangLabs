package main

import (
	"fmt"

)

func main(){
	fmt.Println("learning maps....")

	var languages = make(map[string]string)
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"


	fmt.Println("Languages",languages);
	///loops are interesting in golang

	for key,value := range languages {
		fmt.Printf("For key %v we have value %v\n",key,value);
	}

	//for only values use comma,ok syntax
	for _,value := range languages {
		fmt.Printf("Language value %v\n",value);
	}

	//for only keys use comma,ok syntax
	for key,_ := range languages {
		fmt.Printf("Language key %v\n",key);
	}


}